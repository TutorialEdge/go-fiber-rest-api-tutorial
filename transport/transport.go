package transport

import (
	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

// Setup - set's up our fiber app and the routes
// returns a pointer to app
func Setup() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}
