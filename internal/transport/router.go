package router

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	eventhandler "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/handler"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Register(app *fiber.App,
	pg *pgxpool.Pool,
	rdb *redis.Client,
	rmq *mq.RabbitMQ,
	cfg config.Config) {

	eventhandler := eventhandler.NewEventHandler(pg, rdb, rmq, cfg)

	// app.Get("/health", eventhandler.Health)

	app.Get("/events", eventhandler.ListEvents)
	
	// app.Get("/events/:eventId/seats", eventhandler.ListSeats)
	// app.Post("/holds/seats", eventhandler.HoldSeats)

}
