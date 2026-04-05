package reservationhandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	reservationsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/service"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IReservationHandler interface {
	CreateReservation(c fiber.Ctx) error
	GetReservation(c fiber.Ctx) error
	CleanupExpiredReservations(c fiber.Ctx) error
}

type ReservationHandler struct {
	reservationService *reservationsvc.ReservationService
}

func NewReservationHandler(
	pg *pgxpool.Pool,
	rdb *redis.Client,
	rmq *mq.RabbitMQ,
	cfg config.Config,
) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationsvc.NewReservationService(pg, rdb, 10),
	}
}
