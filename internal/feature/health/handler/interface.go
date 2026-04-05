package healthhandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type IHealthHandler interface {
	Health(c fiber.Ctx) error
}

type HealthHandler struct {
}

func NewHealthHandler(
	pg *pgxpool.Pool,
	rdb *redis.Client,
	rmq *mq.RabbitMQ,
	cfg config.Config) IHealthHandler {

	return &HealthHandler{}

}
