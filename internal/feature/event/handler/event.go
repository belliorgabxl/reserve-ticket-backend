package eventhandler

import (
	"context"

	"github.com/belliorgabxl/reserve-ticket-backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VenueName string `json:"venueName"`
	EventDate string `json:"eventDate"`
}

func (h *EventHandler) ListEvents(c fiber.Ctx) error {
	res, err := h.eventService.ListEvents(context.Background())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, res)
}
