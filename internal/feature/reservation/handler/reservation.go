package reservationhandler

import (
	"context"
	"errors"
	"strconv"

	reservationmodel "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/model"
	reservationsvc "github.com/belliorgabxl/reserve-ticket-backend/internal/feature/reservation/service"
	"github.com/belliorgabxl/reserve-ticket-backend/pkg/response"
	"github.com/gofiber/fiber/v3"
)

func (h *ReservationHandler) CreateReservation(c fiber.Ctx) error {
	var req reservationmodel.CreateReservationRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	res, err := h.reservationService.CreateReservation(context.Background(), req)
	if err != nil {
		switch {
		case errors.Is(err, reservationsvc.ErrInvalidRequest):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, reservationsvc.ErrSeatNotFound):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, reservationsvc.ErrSeatNotHeldByUser):
			return response.Error(c, fiber.StatusConflict, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, res)
}

func (h *ReservationHandler) GetReservation(c fiber.Ctx) error {
	reservationID := c.Params("id")

	res, err := h.reservationService.GetReservation(context.Background(), reservationID)
	if err != nil {
		switch {
		case errors.Is(err, reservationsvc.ErrInvalidRequest):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, reservationsvc.ErrReservationNotFound):
			return response.Error(c, fiber.StatusNotFound, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, err.Error())
		}
	}

	return response.Success(c, res)
}

func (h *ReservationHandler) CleanupExpiredReservations(c fiber.Ctx) error {
	limitStr := c.Query("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	count, err := h.reservationService.CleanupExpiredReservations(context.Background(), limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.Map{
		"expiredReservations": count,
	})
}
