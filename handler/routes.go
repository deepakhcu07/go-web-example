package handler

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {

	app.Get("/", HandleViewHome)
	app.Get("/epoch-converter", HandleViewEpochConverter)
}
