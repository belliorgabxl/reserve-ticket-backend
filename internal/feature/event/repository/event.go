package eventrepository

import (
	"context"

	eventmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/event/model"
)

func (r *EventRepository) ListEvents(ctx context.Context) ([]eventmodel.EventResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT e.id::text, e.name, e.venue_name, e.event_date::text
		FROM events AS e
		ORDER BY e.event_date ASC
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

func (r *EventRepository) ListShowTimesByEventID(ctx context.Context, eventID string) ([]eventmodel.ShowTimeResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id::text, event_id::text, starts_at::text
		FROM show_times
		WHERE event_id = $1
		ORDER BY starts_at ASC
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]eventmodel.ShowTimeResponse, 0)
	for rows.Next() {
		var item eventmodel.ShowTimeResponse
		if err := rows.Scan(
			&item.ID,
			&item.EventID,
			&item.StartAt,
		); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, rows.Err()
}

func (r *EventRepository) GetEventByID(ctx context.Context, eventID string) (*eventmodel.EventResponse, error) {
	row := r.db.QueryRow(ctx, `
		SELECT e.id::text, e.name, e.venue_name, e.event_date::text
		FROM events AS e
		WHERE e.id = $1
		LIMIT 1
	`, eventID)

	var item eventmodel.EventResponse
	if err := row.Scan(
		&item.ID,
		&item.Name,
		&item.VenueName,
		&item.EventDate,
	); err != nil {
		return nil, err
	}

	return &item, nil
}