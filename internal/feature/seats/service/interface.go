package seatsvc

import (
	seatrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/repository"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ISeatService interface {
}

type SeatService struct {
	seatRepository *seatrepository.SeatRepository
	redis          *redis.Client
}

func NewSeatService(
	seatRepository *seatrepository.SeatRepository,
	redis *redis.Client,
) *SeatService {
	return &SeatService{
		seatRepository: seatRepository,
		redis:          redis,
	}
}
