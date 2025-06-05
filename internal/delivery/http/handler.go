package http

import (
	"os"
	"urlshortener/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase *usecase.ShortenerUseCase
}

func NewHandler(u *usecase.ShortenerUseCase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) ShortenURL(c *fiber.Ctx) error {
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

	baseUrl := os.Getenv("BASE_URL")

	return c.JSON(fiber.Map{"short": baseUrl + code})
}

func (h *Handler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")

	url, err := h.usecase.Resolve(code)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "not found")
	}

	return c.Redirect(url, fiber.StatusFound)
}
