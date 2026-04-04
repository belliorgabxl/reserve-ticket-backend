package eventhandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"

)

type IEventHandler interface { 
	ListEvents(c fiber.Ctx) error
}


type EventHandler struct {
	pg  *pgxpool.Pool
	rdb *redis.Client
	rmq *mq.RabbitMQ
	cfg config.Config
}

func NewEventHandler(pg *pgxpool.Pool, rdb *redis.Client, rmq *mq.RabbitMQ, cfg config.Config) IEventHandler {
	return &EventHandler{
		pg:  pg,
		rdb: rdb,
		rmq: rmq,
		cfg: cfg,
	}
}
