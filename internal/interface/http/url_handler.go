package http

import (
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	usecase *usecase.URLUseCase
	baseUrl string
}

func NewURLHandler(u *usecase.URLUseCase, baseUrl string) *URLHandler {
	return &URLHandler{usecase: u, baseUrl: baseUrl}
}

func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var body struct {
		URL string `json:"url"`
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid input")
	}

	code, err := h.usecase.Shorten(body.URL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"short": h.baseUrl + code})
}

func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")

	url, err := h.usecase.Resolve(code)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	}

	return c.Redirect(url, fiber.StatusFound)
}
