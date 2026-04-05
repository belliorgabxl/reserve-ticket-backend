package seatsvc

import (
	"context"
	"errors"

	seatmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/model"
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

	return s.seatRepository.GetSeatsByShowTimeID(ctx, showTimeID)
}