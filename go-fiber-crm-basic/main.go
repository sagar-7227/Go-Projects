package main

import (
	"fmt"
	"github.com/sagar-7227/go-fiber-crm-basic/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sagar-7227/go-fiber-crm-basic/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	database.ConnectDatabase()
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(":8000")
}