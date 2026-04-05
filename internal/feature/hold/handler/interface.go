package holdhandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"

	// redisx "github.com/belliorgabxl/reserve-ticket-backend/pkg/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IHoldHandler interface {
	HoldSeats(c fiber.Ctx) error
}

type HoldHandler struct {
	pg  *pgxpool.Pool
	rdb *redis.Client
	rmq *mq.RabbitMQ
	cfg config.Config
}

func NewHoldHandler(
	pg *pgxpool.Pool,
	rdb *redis.Client,
	rmq *mq.RabbitMQ,
	cfg config.Config,
) *HoldHandler {
	return &HoldHandler{
		pg:  pg,
		rdb: rdb,
		rmq: rmq,
		cfg: cfg,
	}
}
