package healthhandler

import (
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VenueName string `json:"venueName"`
	EventDate string `json:"eventDate"`
}

func (h *HealthHandler) Health(c fiber.Ctx) error {
	return response.Success(c, fiber.Map{
		"status": "ok",
	})
}
