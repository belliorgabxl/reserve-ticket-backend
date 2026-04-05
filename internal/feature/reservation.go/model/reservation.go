package reservationmodel

import "time"

type ReservationStatus string

const (
	ReservationStatusHolding   ReservationStatus = "HOLDING"
	ReservationStatusExpired   ReservationStatus = "EXPIRED"
	ReservationStatusConfirmed ReservationStatus = "CONFIRMED"
	ReservationStatusCancelled ReservationStatus = "CANCELLED"
)

type CreateReservationRequest struct {
	UserID     string   `json:"userId"`
	ShowTimeID string   `json:"showTimeId"`
	SeatIDs    []string `json:"seatIds"`
}

type CreateReservationResponse struct {
	ReservationID string    `json:"reservationId"`
	Status        string    `json:"status"`
	ExpiresAt     time.Time `json:"expiresAt"`
	SeatIDs       []string  `json:"seatIds"`
}

type Reservation struct {
	ID         string            `json:"id"`
	UserID     string            `json:"userId"`
	ShowTimeID string            `json:"showTimeId"`
	Status     ReservationStatus `json:"status"`
	ExpiresAt  time.Time         `json:"expiresAt"`
	CreatedAt  time.Time         `json:"createdAt"`
	Items      []ReservationItem `json:"items"`
}

type ReservationItem struct {
	ID            string    `json:"id"`
	ReservationID string    `json:"reservationId"`
	SeatID        string    `json:"seatId"`
	SeatCode      string    `json:"seatCode"`
	ZoneID        string    `json:"zoneId"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"createdAt"`
}

type SeatInfo struct {
	SeatID   string
	SeatCode string
	ZoneID   string
	Price    float64
}
