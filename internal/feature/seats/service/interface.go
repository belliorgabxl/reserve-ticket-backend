package seatsvc

import (
	seatrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ISeatService interface {
}

type SeatService struct {
	redis          *redis.Client
	holdTTLMinutes int
	seatRepository seatrepository.SeatRepository
}

func NewSeatService(pg *pgxpool.Pool, redis *redis.Client,holdTTLMinutes int) *SeatService {
	return &SeatService{
		seatRepository: *seatrepository.NewSeatRepository(pg),
	}

}
