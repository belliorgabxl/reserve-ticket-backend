package reservationhandler

import (
	reservationsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/service"
	"github.com/gofiber/fiber/v3"
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
	reservationService *reservationsvc.ReservationService,
) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}