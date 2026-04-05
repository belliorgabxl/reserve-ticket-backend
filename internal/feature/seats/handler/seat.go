package seathandler

import (
	"context"
	// "errors"
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

func (h *SeatHandler) GetSeatsByShowTimeID(c fiber.Ctx) error {
	showTimeID := c.Params("showTimeId")

	res, err := h.seatService.GetSeatsByShowTimeID(context.Background(), showTimeID)
	if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, res)
}