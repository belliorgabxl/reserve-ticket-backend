package bookinghandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IBookingHandler interface {
	HoldSeats(c fiber.Ctx) error
}

type BookingHandler struct {
	pg  *pgxpool.Pool
	rdb *redis.Client
	rmq *mq.RabbitMQ
	cfg config.Config
}

func NewBookingHandler(pg *pgxpool.Pool, rdb *redis.Client, rmq *mq.RabbitMQ, cfg config.Config) IBookingHandler {
	return &BookingHandler{
		pg:  pg,
		rdb: rdb,
		rmq: rmq,
		cfg: cfg,
	}
}
