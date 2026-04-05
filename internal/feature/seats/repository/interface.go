package seatrepository

import (
	"context"

	seatmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatRepository struct {
	db *pgxpool.Pool
}

func NewSeatRepository(db *pgxpool.Pool) *SeatRepository {
	return &SeatRepository{db: db}
}

type ISeatRepository interface {
	GetSeatsByShowTimeID(ctx context.Context, showTimeID string) ([]seatmodel.SeatItem, error)
}