package url

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	baseUrl = "http://localhost:8888"
)

type UrlHandler struct {
	urlUsecase *UrlUsecase
}

func NewUrlHandler(urlUsecase *UrlUsecase) *UrlHandler {
	return &UrlHandler{urlUsecase: urlUsecase}
}

func (u *UrlHandler) Shortener(c *fiber.Ctx) error {
	var input CreateUrl

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	url, new, err := u.urlUsecase.SearchOrCreateNewUrl(input.Original)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var status int
	if new {
		status = fiber.StatusCreated
	} else {
		status = fiber.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", baseUrl, url.ID)

	return c.Status(status).JSON(fiber.Map{
		"Location": shortUrl,
		"Url":      url,
	})
}

func (u *UrlHandler) Redirect(c *fiber.Ctx) error {
	id := c.Params("id")
	url, b := u.urlUsecase.Find(id)

	if b {
		return c.Redirect(url.Original, fiber.StatusMovedPermanently)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
		"Url":     url,
	})
}
