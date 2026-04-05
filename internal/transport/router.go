package router

import (
	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	eventhandler "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/handler"
	eventrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/repository"
	eventsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/service"
	healthhandler "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/health/handler"
	holdhandler "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/handler"
	holdsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/service"
	reservationhandler "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/handler"
	reservationrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/repository"
	reservationsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/service"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Register(
	app *fiber.App,
	pg *pgxpool.Pool,
	rdb *redis.Client,
	rmq *mq.RabbitMQ,
	cfg config.Config,
) {
	health := healthhandler.NewHealthHandler(pg, rdb, rmq, cfg)

	holdService := holdsvc.NewHoldService(rdb, cfg.HoldTTLMinutes)
	holdHandler := holdhandler.NewHoldHandler(holdService)

	eventRepo := eventrepository.NewEventRepository(pg)
	eventService := eventsvc.NewEventService(eventRepo)
	eventHandler := eventhandler.NewEventHandler(eventService)

	reservationRepo := reservationrepository.NewReservationRepository(pg)
	reservationService := reservationsvc.NewReservationService(reservationRepo, rdb, cfg.HoldTTLMinutes)
	reservationHandler := reservationhandler.NewReservationHandler(reservationService)

	app.Get("/health", health.Health)
	app.Get("/events", eventHandler.ListEvents)
	app.Post("/holds/seats", holdHandler.HoldSeats)

	app.Post("/reservations", reservationHandler.CreateReservation)
	app.Get("/reservations/:id", reservationHandler.GetReservation)
	app.Post("/internal/reservations/cleanup-expired", reservationHandler.CleanupExpiredReservations)
}
