package main

import (
	"log"
	"wedding-api/config"
	"wedding-api/model"
	"wedding-api/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	config.ConnectDB()

	db := config.DB
	// migration
	err := db.AutoMigrate(
		&model.City{},
		&model.WeddingOrganizer{},
		&model.WeddingPackage{},
		&model.BonusPackage{},
		&model.WeddingBonusPackage{},
		&model.WeddingPhoto{},
		&model.WeddingTestimonial{},
		&model.BookingTransaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
