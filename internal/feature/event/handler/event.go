package eventhandler

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

type EventResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VenueName string `json:"venueName"`
	EventDate string `json:"eventDate"`
}

func (h *EventHandler) ListEvents(c fiber.Ctx) error {
	rows, err := h.pg.Query(context.Background(), `
		SELECT id::text, name, venue_name, event_date::text
		FROM events
		ORDER BY event_date ASC
		LIMIT 20
	`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer rows.Close()

	out := make([]EventResponse, 0)
	for rows.Next() {
		var item EventResponse
		if err := rows.Scan(&item.ID, &item.Name, &item.VenueName, &item.EventDate); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		out = append(out, item)
	}

	return c.JSON(out)
}
