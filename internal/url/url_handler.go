package url

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	baseUrl = "http://localhost:8888"
)

type urlHandler struct {
	urlUsecase UrlUsecase
}

func NewUrlHandler(urlUsecase UrlUsecase) *urlHandler {
	return &urlHandler{urlUsecase: urlUsecase}
}

func (u *urlHandler) Shortener(c *fiber.Ctx) error {
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

func (u *urlHandler) Redirect(c *fiber.Ctx) error {
	id := c.Params("id")
	url, exist := u.urlUsecase.Find(id)

	if exist {
		return c.Redirect(url.Original, fiber.StatusMovedPermanently)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
		"Url":     url,
	})
}

func (u *urlHandler) Status(c *fiber.Ctx) error {
	id := c.Params("id")
	url, exist := u.urlUsecase.Status(id)

	if exist {
		return c.Status(fiber.StatusFound).JSON(url)
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
		"Url":     url,
	})
}
