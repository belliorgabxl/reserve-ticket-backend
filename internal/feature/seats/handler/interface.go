package seathandler

import (
	// seatsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seat/service"
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	seatsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/service"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type SeatHandler struct {
	seatService *seatsvc.SeatService
}

type ISeatHandler interface {
	GetSeatsByShowTimeID(c fiber.Ctx) error
}

func NewSeatHandler(pg *pgxpool.Pool, rdb *redis.Client,
	cfg config.Config) *SeatHandler {
	return &SeatHandler{
		seatService: seatsvc.NewSeatService(pg, rdb, 10),
	}
}
