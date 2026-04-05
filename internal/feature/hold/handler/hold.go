package holdhandler

import (
	holdmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/model"
	"github.com/gofiber/fiber/v3"
	"context"
	"time"
	"fmt"
)
func (h *HoldHandler) HoldSeats(c fiber.Ctx) error {
	var req holdmodel.HoldSeatsRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
		})
	}

	if req.ShowTimeID == "" || len(req.SeatIDs) == 0 || req.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "showTimeId, seatIds, userId are required",
		})
	}

	ctx := context.Background()
	ttl := 10 * time.Minute

	held := make([]string, 0, len(req.SeatIDs))
	failed := make([]string, 0)

	for _, seatID := range req.SeatIDs {
		key := fmt.Sprintf("seat:hold:%s:%s", req.ShowTimeID, seatID)

		ok, err := h.rdb.SetNX(ctx, key, req.UserID, ttl).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if ok {
			held = append(held, seatID)
		} else {
			failed = append(failed, seatID)
		}
	}

	return c.JSON(fiber.Map{
		"heldSeatIds":   held,
		"failedSeatIds": failed,
		"expiresInSec":  int(ttl.Seconds()),
	})
}