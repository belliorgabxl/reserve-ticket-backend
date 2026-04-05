package reservationsvc

import (
	"context"
	"errors"
	"fmt"
	"time"

	reservationmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/model"
	//  "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/repository"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidRequest      = errors.New("invalid request")
	ErrSeatNotFound        = errors.New("seat not found")
	ErrSeatNotHeldByUser   = errors.New("one or more seats are not held by this user")
	ErrReservationNotFound = errors.New("reservation not found")
)

func (s *ReservationService) CreateReservation(
	ctx context.Context,
	req reservationmodel.CreateReservationRequest,
) (*reservationmodel.CreateReservationResponse, error) {
	if req.UserID == "" || req.ShowTimeID == "" || len(req.SeatIDs) == 0 {
		return nil, ErrInvalidRequest
	}

	seatMap := make(map[string]struct{}, len(req.SeatIDs))
	for _, seatID := range req.SeatIDs {
		if seatID == "" {
			return nil, ErrInvalidRequest
		}
		seatMap[seatID] = struct{}{}
	}

	seats, err := s.reservationRepo.FindSeatsByIDs(ctx, req.ShowTimeID, req.SeatIDs)
	if err != nil {
		return nil, err
	}
	if len(seats) != len(seatMap) {
		return nil, ErrSeatNotFound
	}

	for _, seatID := range req.SeatIDs {
		key := s.seatHoldKey(req.ShowTimeID, seatID)

		val, err := s.redis.Get(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return nil, ErrSeatNotHeldByUser
			}
			return nil, err
		}

		if val != req.UserID {
			return nil, ErrSeatNotHeldByUser
		}
	}

	ttl := time.Duration(s.holdTTLMinutes) * time.Minute
	expiresAt := time.Now().Add(ttl)

	reservation, err := s.reservationRepo.CreateReservation(ctx, req.UserID, req.ShowTimeID, expiresAt, seats)
	if err != nil {
		return nil, err
	}

	seatIDs := make([]string, 0, len(reservation.Items))
	for _, item := range reservation.Items {
		seatIDs = append(seatIDs, item.SeatID)
	}

	return &reservationmodel.CreateReservationResponse{
		ReservationID: reservation.ID,
		Status:        string(reservation.Status),
		ExpiresAt:     reservation.ExpiresAt,
		SeatIDs:       seatIDs,
	}, nil
}

func (s *ReservationService) GetReservation(
	ctx context.Context,
	reservationID string,
) (*reservationmodel.Reservation, error) {
	if reservationID == "" {
		return nil, ErrInvalidRequest
	}

	res, err := s.reservationRepo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, ErrReservationNotFound
	}

	return res, nil
}

func (s *ReservationService) CleanupExpiredReservations(
	ctx context.Context,
	limit int,
) (int, error) {
	if limit <= 0 {
		limit = 100
	}

	reservations, err := s.reservationRepo.FindExpiredHoldingReservations(ctx, limit)
	if err != nil {
		return 0, err
	}

	count := 0

	for _, reservation := range reservations {
		seatIDs, err := s.reservationRepo.GetReservationSeatIDs(ctx, reservation.ID)
		if err != nil {
			return count, err
		}

		if err := s.reservationRepo.MarkReservationExpired(ctx, reservation.ID); err != nil {
			return count, err
		}

		for _, seatID := range seatIDs {
			key := s.seatHoldKey(reservation.ShowTimeID, seatID)
			_ = s.redis.Del(ctx, key).Err()
		}

		count++
	}

	return count, nil
}

func (s *ReservationService) seatHoldKey(showTimeID, seatID string) string {
	return fmt.Sprintf("seat:hold:%s:%s", showTimeID, seatID)
}
