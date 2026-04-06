package main

import (
	"log"

	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	router "github.com/belliorgabxl/reserve-ticket-backend/internal/transport"
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/database"
	mq "github.com/belliorgabxl/reserve-ticket-backend/pkg/rabbitmq"
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/redis"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.MustLoad()

	pg, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := redisx.NewRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	rmq, err := mq.NewRabbitMQ(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer rmq.Close()

	app := fiber.New()
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
	}))

	router.Register(app, pg, rdb, rmq, cfg)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
