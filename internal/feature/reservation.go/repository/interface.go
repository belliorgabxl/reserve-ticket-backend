package reservationrepository

import (
	"context"
	"time"

	reservationmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IReservationRepository interface {
	GetReservationByID(ctx context.Context, reservationID string) (*reservationmodel.Reservation, error)
	FindSeatsByIDs(ctx context.Context, showTimeID string, seatIDs []string) ([]reservationmodel.SeatInfo, error)
	CreateReservation(ctx context.Context, userID string, showTimeID string, expiresAt time.Time, seats []reservationmodel.SeatInfo) (*reservationmodel.Reservation, error)
	FindExpiredHoldingReservations(ctx context.Context, limit int) ([]reservationmodel.Reservation, error)
	GetReservationSeatIDs(ctx context.Context, reservationID string) ([]string, error)
	MarkReservationExpired(ctx context.Context, reservationID string) error
}


type ReservationRepository struct {
	db *pgxpool.Pool
}

func NewReservationRepository(db *pgxpool.Pool) *ReservationRepository {
	return &ReservationRepository{db: db}
}

