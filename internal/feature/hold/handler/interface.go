package holdhandler

import (
	holdsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/service"
	"github.com/gofiber/fiber/v3"
)

type IHoldHandler interface {
	HoldSeats(c fiber.Ctx) error
	ReleaseSeats(c fiber.Ctx) error
}

type HoldHandler struct {
	holdService *holdsvc.HoldService
}

func NewHoldHandler(holdService *holdsvc.HoldService) *HoldHandler {
	return &HoldHandler{
		holdService: holdService,
	}
}
