package reservationsvc

import (
	"context"

	reservationmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/model"
	reservationrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IReservationService interface {
	CreateReservation(ctx context.Context, req reservationmodel.CreateReservationRequest) (*reservationmodel.CreateReservationResponse, error)
	GetReservation(ctx context.Context, reservationID string) (*reservationmodel.Reservation, error)
	CleanupExpiredReservations(ctx context.Context, limit int) (int, error)
}

type ReservationService struct {
	reservationRepo *reservationrepository.ReservationRepository
	redis           *redis.Client
	holdTTLMinutes  int
}

func NewReservationService(
	pg *pgxpool.Pool,
	redis *redis.Client,
	holdTTLMinutes int,
) *ReservationService {
	return &ReservationService{
		reservationRepo: reservationrepository.NewReservationRepository(pg),
		redis:           redis,
		holdTTLMinutes:  holdTTLMinutes,
	}
}