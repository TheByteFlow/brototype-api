package main

import (
	"log"

	"github.com/brototype-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/auth", handlers.GetUserAuth)
	api.Get("/details", handlers.GetUserDetails)
	api.Get("/reviews", handlers.GetUserReviews)
	api.Get("/foundations", handlers.GetFoundationReviews)
}
func main() {
	app := fiber.New()
	setupRoutes(app)
	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
