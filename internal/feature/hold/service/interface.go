package holdsvc

import (
	"context"

	holdmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/model"
	"github.com/redis/go-redis/v9"
)

type HoldService struct {
	redis          *redis.Client
	holdTTLMinutes int
}

type IHoldService interface {
	HoldSeats(ctx context.Context, req holdmodel.HoldSeatsRequest) (*holdmodel.HoldSeatsResponse, error)
	ReleaseSeats(ctx context.Context, req holdmodel.ReleaseSeatsRequest) error
}

func NewHoldService(redis *redis.Client, holdTTLMinutes int) *HoldService {
	return &HoldService{
		redis:          redis,
		holdTTLMinutes: holdTTLMinutes,
	}
}
