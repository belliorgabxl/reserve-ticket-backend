package seathandler

import (
	seatsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/service"
	"github.com/gofiber/fiber/v3"
)

type SeatHandler struct {
	seatService *seatsvc.SeatService
}

type ISeatHandler interface {
	GetSeatsByShowTimeID(c fiber.Ctx) error
}

func NewSeatHandler(seatService *seatsvc.SeatService) *SeatHandler {
	return &SeatHandler{
		seatService: seatService,
	}
}