package eventrepository

import (
	"context"

	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
)

func (r *EventRepository) ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id::text, name, venue_name, event_date::text
		FROM events
		ORDER BY event_date ASC
		LIMIT 20
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]eventmodel.EventResponse, 0)
	for rows.Next() {
		var item eventmodel.EventResponse
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.VenueName,
			&item.EventDate,
		); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, rows.Err()
}