package holdhandler

import (
	"context"
	"errors"

	holdmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/model"
	holdsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/hold/service"
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

func (h *HoldHandler) HoldSeats(c fiber.Ctx) error {
	var req holdmodel.HoldSeatsRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	res, err := h.holdService.HoldSeats(context.Background(), req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())

	}

	return response.Success(c, res)
}

func (h *HoldHandler) ReleaseSeats(c fiber.Ctx) error {
	var req holdmodel.ReleaseSeatsRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	err := h.holdService.ReleaseSeats(context.Background(), req)
	if err != nil {
		switch {
		case errors.Is(err, holdsvc.ErrInvalidRequest):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, fiber.Map{
		"released": true,
	})
}
