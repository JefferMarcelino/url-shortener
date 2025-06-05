package http

import (
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AnalyticsHandler struct {
	usecase *usecase.AnalyticsUseCase
}

func NewAnalyticsHandler(usecase *usecase.AnalyticsUseCase) *AnalyticsHandler {
	return &AnalyticsHandler{usecase: usecase}
}

func (h *AnalyticsHandler) GetAnalyticsByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	analytics, err := h.usecase.GetAnalyticsByCode(code)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "an unexpected error occured")
	}

	return c.JSON(analytics)
}
