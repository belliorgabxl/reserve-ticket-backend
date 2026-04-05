package seatrepository

import (
	"context"

	// "github.com/belliorgabxl/reserve-ticket-backend/internal/entities"
	seatmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/seats/model"
)

func (r *SeatRepository) GetSeatsByShowTimeID(
	ctx context.Context,
	showTimeID string,
) ([]seatmodel.SeatItem, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id::text,
			show_time_id::text,
			zone_id::text,
			seat_code,
			COALESCE(row_label, ''),
			COALESCE(seat_number, 0),
			COALESCE(price, 0),
			status
		FROM seats
		WHERE show_time_id = $1
		ORDER BY seat_code ASC
	`, showTimeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]seatmodel.SeatItem, 0)
	for rows.Next() {
		var item seatmodel.SeatItem
		if err := rows.Scan(
			&item.SeatID,
			&item.ShowTimeID,
			&item.ZoneID,
			&item.SeatCode,
			&item.RowLabel,
			&item.SeatNumber,
			&item.Price,
			&item.Status,
		); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, rows.Err()
}