package analytics

import (
	"urlshortener/internal/application"

	"github.com/gofiber/fiber/v2"
)

type AnalyticsHandler struct {
	usecase *application.AnalyticsUseCase
}

func NewAnalyticsHandler(usecase *application.AnalyticsUseCase) *AnalyticsHandler {
	return &AnalyticsHandler{usecase: usecase}
}

func (h *AnalyticsHandler) RegisterAnalyticsRoutes(app *fiber.App) {
	app.Get("/analytics/:code", h.GetClickEventsByCode)
}

func (h *AnalyticsHandler) GetClickEventsByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	clickEvents, err := h.usecase.GetClickEventsByCode(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "analytics unavailable"})
	}

	dtoEvents := make([]ClickEventDTO, 0, len(clickEvents))
	for _, e := range clickEvents {
		dtoEvents = append(dtoEvents, ClickEventDTO{
			Timestamp: e.Timestamp,
			IP:        e.IP,
			UserAgent: e.UserAgent,
		})
	}

	return c.JSON(AnalyticsResponse{
		ShortCode:   code,
		TotalClicks: len(dtoEvents),
		Clicks:      dtoEvents,
	})
}
