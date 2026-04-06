package seatsvc

import (
	"context"
	"errors"

	seatmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/model"
	"fmt"
	// seatmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seat/model"
)

var ErrInvalidShowTimeID = errors.New("invalid showTimeId")

func (s *SeatService) GetSeatsByShowTimeID(
	ctx context.Context,
	showTimeID string,
) ([]seatmodel.SeatItem, error) {
	if showTimeID == "" {
		return nil, ErrInvalidShowTimeID
	}

	seats, err := s.seatRepository.GetSeatsByShowTimeID(ctx, showTimeID)
	if err != nil {
		return nil, err
	}

	for i := range seats {
		// ถ้า booked ใน DB อยู่แล้ว ไม่ต้องเช็ก hold
		if seats[i].Status == seatmodel.SeatStatusBooked {
			continue
		}

		key := fmt.Sprintf("seat:hold:%s:%s", showTimeID, seats[i].SeatID)

		val, err := s.redis.Get(ctx, key).Result()
		if err != nil {
			// ไม่มี key = ยังไม่ถูก hold
			// redis.Nil ไม่ต้อง error
			continue
		}

		ttl, ttlErr := s.redis.TTL(ctx, key).Result()
		if ttlErr == nil && ttl > 0 {
			seats[i].HoldExpiresIn = int64(ttl.Seconds())
		}

		if val != "" {
			seats[i].Status = seatmodel.SeatStatusHeld
			// seats[i].HeldByUserID = val
		}
	}

	return seats, nil
}