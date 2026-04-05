package seatmodel


type SeatStatus string

const (
	SeatStatusAvailable SeatStatus = "AVAILABLE"
	SeatStatusHeld      SeatStatus = "HELD"
	SeatStatusBooked    SeatStatus = "BOOKED"
)

type SeatItem struct {
	SeatID     string     `json:"seatId"`
	ShowTimeID string     `json:"showTimeId"`
	ZoneID     string     `json:"zoneId"`
	SeatCode   string     `json:"seatCode"`
	RowLabel   string     `json:"rowLabel"`
	SeatNumber int        `json:"seatNumber"`
	Price      float64    `json:"price"`
	Status     SeatStatus `json:"status"`
}