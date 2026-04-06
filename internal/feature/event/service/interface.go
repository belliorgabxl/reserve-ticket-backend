package eventsvc

import (
	"context"
	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
	eventrepository "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/repository"
)

type IEventService interface {
	ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error)
	ListShowTimesByEventID(ctx context.Context, eventId string) ([]eventmodel.ShowTimeResponse, error)
}

type EventService struct {
	eventRepo *eventrepository.EventRepository
}

func NewEventService(eventRepo *eventrepository.EventRepository) *EventService {
	return &EventService{
		eventRepo: eventRepo,
	}
}
