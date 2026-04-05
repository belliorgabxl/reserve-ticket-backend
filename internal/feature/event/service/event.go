package eventsvc

import (
	"context"

	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
)

func (s *EventService) ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error) {
	return s.eventRepo.ListEvents(ctx)
}