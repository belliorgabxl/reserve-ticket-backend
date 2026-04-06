package eventsvc

import (
	"context"

	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
)

func (s *EventService) ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error) {
	return s.eventRepo.ListEvents(ctx)
}

func (s *EventService) ListShowTimesByEventID(ctx context.Context, eventId string) ([]eventmodel.ShowTimeResponse, error) {
	return s.eventRepo.ListShowTimesByEventID(ctx, eventId)
}
func (s *EventService) GetEventByID(ctx context.Context, eventId string) (*eventmodel.EventResponse, error) {
	return s.eventRepo.GetEventByID(ctx, eventId)
}
