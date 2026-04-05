package reservationrepository

import (
	"context"
	"errors"
	"time"

	reservationmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation.go/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ReservationRepository) FindSeatsByIDs(
	ctx context.Context,
	showTimeID string,
	seatIDs []string,
) ([]reservationmodel.SeatInfo, error) {
	rows, err := r.db.Query(ctx, `
		SELECT s.id::text, s.seat_code, s.zone_id::text, COALESCE(s.price, 0)
		FROM seats s
		WHERE s.show_time_id = $1
		  AND s.id = ANY($2)
	`, showTimeID, seatIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]reservationmodel.SeatInfo, 0, len(seatIDs))
	for rows.Next() {
		var item reservationmodel.SeatInfo
		if err := rows.Scan(&item.SeatID, &item.SeatCode, &item.ZoneID, &item.Price); err != nil {
			return nil, err
		}
		out = append(out, item)
	}

	return out, rows.Err()
}

func (r *ReservationRepository) CreateReservation(
	ctx context.Context,
	userID string,
	showTimeID string,
	expiresAt time.Time,
	seats []reservationmodel.SeatInfo,
) (*reservationmodel.Reservation, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	reservationID := uuid.NewString()
	createdAt := time.Now()

	_, err = tx.Exec(ctx, `
		INSERT INTO reservations (
			id, user_id, show_time_id, status, expires_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`,
		reservationID,
		userID,
		showTimeID,
		reservationmodel.ReservationStatusHolding,
		expiresAt,
		createdAt,
	)
	if err != nil {
		return nil, err
	}

	items := make([]reservationmodel.ReservationItem, 0, len(seats))
	for _, seat := range seats {
		itemID := uuid.NewString()
		itemCreatedAt := time.Now()

		_, err = tx.Exec(ctx, `
			INSERT INTO reservation_items (
				id, reservation_id, seat_id, price, created_at
			) VALUES ($1, $2, $3, $4, $5)
		`,
			itemID,
			reservationID,
			seat.SeatID,
			seat.Price,
			itemCreatedAt,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, reservationmodel.ReservationItem{
			ID:            itemID,
			ReservationID: reservationID,
			SeatID:        seat.SeatID,
			SeatCode:      seat.SeatCode,
			ZoneID:        seat.ZoneID,
			Price:         seat.Price,
			CreatedAt:     itemCreatedAt,
		})
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &reservationmodel.Reservation{
		ID:         reservationID,
		UserID:     userID,
		ShowTimeID: showTimeID,
		Status:     reservationmodel.ReservationStatusHolding,
		ExpiresAt:  expiresAt,
		CreatedAt:  createdAt,
		Items:      items,
	}, nil
}

func (r *ReservationRepository) GetReservationByID(
	ctx context.Context,
	reservationID string,
) (*reservationmodel.Reservation, error) {
	var res reservationmodel.Reservation

	err := r.db.QueryRow(ctx, `
		SELECT id::text, user_id, show_time_id::text, status, expires_at, created_at
		FROM reservations
		WHERE id = $1
	`, reservationID).Scan(
		&res.ID,
		&res.UserID,
		&res.ShowTimeID,
		&res.Status,
		&res.ExpiresAt,
		&res.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	rows, err := r.db.Query(ctx, `
		SELECT 
			ri.id::text,
			ri.reservation_id::text,
			ri.seat_id::text,
			s.seat_code,
			s.zone_id::text,
			ri.price,
			ri.created_at
		FROM reservation_items ri
		JOIN seats s ON s.id = ri.seat_id
		WHERE ri.reservation_id = $1
		ORDER BY s.seat_code ASC
	`, reservationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]reservationmodel.ReservationItem, 0)
	for rows.Next() {
		var item reservationmodel.ReservationItem
		if err := rows.Scan(
			&item.ID,
			&item.ReservationID,
			&item.SeatID,
			&item.SeatCode,
			&item.ZoneID,
			&item.Price,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	res.Items = items
	return &res, rows.Err()
}

func (r *ReservationRepository) FindExpiredHoldingReservations(
	ctx context.Context,
	limit int,
) ([]reservationmodel.Reservation, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id::text, user_id, show_time_id::text, status, expires_at, created_at
		FROM reservations
		WHERE status = $1
		  AND expires_at < NOW()
		ORDER BY expires_at ASC
		LIMIT $2
	`, reservationmodel.ReservationStatusHolding, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]reservationmodel.Reservation, 0)
	for rows.Next() {
		var res reservationmodel.Reservation
		if err := rows.Scan(
			&res.ID,
			&res.UserID,
			&res.ShowTimeID,
			&res.Status,
			&res.ExpiresAt,
			&res.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, res)
	}

	return out, rows.Err()
}

func (r *ReservationRepository) GetReservationSeatIDs(
	ctx context.Context,
	reservationID string,
) ([]string, error) {
	rows, err := r.db.Query(ctx, `
		SELECT seat_id::text
		FROM reservation_items
		WHERE reservation_id = $1
	`, reservationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]string, 0)
	for rows.Next() {
		var seatID string
		if err := rows.Scan(&seatID); err != nil {
			return nil, err
		}
		out = append(out, seatID)
	}
	return out, rows.Err()
}

func (r *ReservationRepository) MarkReservationExpired(
	ctx context.Context,
	reservationID string,
) error {
	cmd, err := r.db.Exec(ctx, `
		UPDATE reservations
		SET status = $1
		WHERE id = $2
		  AND status = $3
	`,
		reservationmodel.ReservationStatusExpired,
		reservationID,
		reservationmodel.ReservationStatusHolding,
	)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return nil
	}

	return nil
}
