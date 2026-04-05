package holdsvc

import (
	"context"
	"errors"
	"fmt"
	"time"

	holdmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/model"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
)

func (s *HoldService) HoldSeats(
	ctx context.Context,
	req holdmodel.HoldSeatsRequest,
) (*holdmodel.HoldSeatsResponse, error) {
	if req.UserID == "" || req.ShowTimeID == "" || len(req.SeatIDs) == 0 {
		return nil, ErrInvalidRequest
	}

	ttl := time.Duration(s.holdTTLMinutes) * time.Minute

	held := make([]string, 0, len(req.SeatIDs))
	failed := make([]string, 0)

	for _, seatID := range req.SeatIDs {
		if seatID == "" {
			failed = append(failed, seatID)
			continue
		}

		key := s.seatHoldKey(req.ShowTimeID, seatID)
		ok, err := s.redis.SetNX(ctx, key, req.UserID, ttl).Result()
		if err != nil {
			return nil, err
		}

		if ok {
			held = append(held, seatID)
		} else {
			failed = append(failed, seatID)
		}
	}

	return &holdmodel.HoldSeatsResponse{
		HeldSeatIDs:   held,
		FailedSeatIDs: failed,
		ExpiresInSec:  int(ttl.Seconds()),
	}, nil
}

func (s *HoldService) ReleaseSeats(
	ctx context.Context,
	req holdmodel.ReleaseSeatsRequest,
) error {
	if req.UserID == "" || req.ShowTimeID == "" || len(req.SeatIDs) == 0 {
		return ErrInvalidRequest
	}

	for _, seatID := range req.SeatIDs {
		key := s.seatHoldKey(req.ShowTimeID, seatID)

		val, err := s.redis.Get(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			return err
		}

		if val == req.UserID {
			if err := s.redis.Del(ctx, key).Err(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *HoldService) seatHoldKey(showTimeID, seatID string) string {
	return fmt.Sprintf("seat:hold:%s:%s", showTimeID, seatID)
}
