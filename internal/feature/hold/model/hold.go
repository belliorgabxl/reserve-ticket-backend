package holdmodel

type HoldSeatsRequest struct {
	UserID     string   `json:"userId"`
	ShowTimeID string   `json:"showTimeId"`
	SeatIDs    []string `json:"seatIds"`
}

type HoldSeatsResponse struct {
	HeldSeatIDs   []string `json:"heldSeatIds"`
	FailedSeatIDs []string `json:"failedSeatIds"`
	ExpiresInSec  int      `json:"expiresInSec"`
	IsPartial     bool     `json:"isPartial"`
}
type ReleaseSeatsRequest struct {
	UserID     string   `json:"userId"`
	ShowTimeID string   `json:"showTimeId"`
	SeatIDs    []string `json:"seatIds"`
}
