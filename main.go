package main

import (
	"fmt"
	"github.com/deepakhcu07/go-web-example/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	fmt.Printf("\n gofiber example")

	engine := html.New("./views", ".html")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base.layout",
	})

	app.Static("/", "./public")

	// Routes
	handler.Setup(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
