package url

import (
	"urlshortener/internal/application"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	usecase *application.URLUseCase
	baseUrl string
}

func NewURLHandler(u *application.URLUseCase, baseUrl string) *URLHandler {
	return &URLHandler{usecase: u, baseUrl: baseUrl}
}

func (h *URLHandler) RegisterURLRoutes(app *fiber.App) {
	app.Post("/shorten", h.ShortenURL)
	app.Get("/:code", h.Redirect)
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var body ShortenRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid input")
	}

	code, err := h.usecase.Shorten(body.OriginalURL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(ShortenResponse{
		ShortCode:    code,
		OriginalURL:  body.OriginalURL,
		ShortURL:     h.baseUrl + code,
		AnalyticsURL: h.baseUrl + "analytics/" + code,
	})
}

func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")
	ip := c.IP()
	ua := c.Get("User-Agent")

	url, err := h.usecase.Resolve(code, ip, ua)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	}

	c.Set("Cache-Control", "no-store")
	return c.Redirect(url, fiber.StatusFound)
}
