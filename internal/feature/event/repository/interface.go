package eventrepository

import (
	"context"

	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IEventRepository interface {
	ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error)
}

type EventRepository struct {
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		db: db,
	}
}