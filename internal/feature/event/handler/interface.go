package eventhandler

import (
	eventsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/service"
	"github.com/gofiber/fiber/v3"
)

type IEventHandler interface {
	ListEvents(c fiber.Ctx) error
}

type EventHandler struct {
	eventService *eventsvc.EventService
}

func NewEventHandler(eventService *eventsvc.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}
