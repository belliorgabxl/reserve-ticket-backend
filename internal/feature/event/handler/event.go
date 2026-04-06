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
func (h *EventHandler) ListShowTimesByEventID(c fiber.Ctx) error {
	eventId := c.Params("eventId")

	res, err := h.eventService.ListShowTimesByEventID(context.Background(), eventId)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, res)
}
func (h *EventHandler) GetEventByID(c fiber.Ctx) error {
	eventId := c.Params("eventId")

	res, err := h.eventService.GetEventByID(context.Background(), eventId)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, res)
}
