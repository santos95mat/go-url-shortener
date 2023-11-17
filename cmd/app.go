package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/santos95mat/go-url-shortener/initializer"
	"github.com/santos95mat/go-url-shortener/internal/url"
)

var (
	app               = fiber.New()
	memory_repository = url.NewRepositoryMemory()
	url_usecase       = url.NewUrlUsecase(memory_repository)
	url_handler       = url.NewUrlHandler(url_usecase)
)

func init() {
	initializer.LoadEnvVariables()
}

func main() {
	app.Use(cors.New())

	app.Post("/api/shorten", url_handler.Shortener)
	app.Get("/r/:id", url_handler.Redirect)
	app.Get("/r/status/:id", url_handler.Status)

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}
