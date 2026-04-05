package eventmodel

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VenueName string `json:"venueName"`
	EventDate string `json:"eventDate"`
}