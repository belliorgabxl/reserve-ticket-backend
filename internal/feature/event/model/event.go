package eventmodel

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VenueName string `json:"venueName"`
	EventDate string `json:"eventDate"`
}

type ShowTimeResponse struct {
	ID        string `json:"id"`
	EventID   string `json:"event_id"`
	StartAt string `json:"starts_at"`
}
