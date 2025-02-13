package main

import (
	"log"

	"github.com/brototype-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
			AllowHeaders: "Origin, Content-Type, Accept",
		},
	))
	setupRoutes(app)
	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
