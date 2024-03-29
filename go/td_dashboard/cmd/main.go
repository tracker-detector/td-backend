package main

import (
	"github.com/Tracking-Detector/td-backend/go/td_dashboard/handlers"
	"github.com/Tracking-Detector/td-backend/go/td_dashboard/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "static")
	statusService := services.NewStatusService(services.NewRestyRestService())
	datasetService := services.NewDatasetService(services.NewRestyRestService())
	handlers.NewHomeHandler(app, statusService).RegisterHandlers()
	handlers.NewDatasetHandler(app, datasetService).RegisterHandlers()
	handlers.NewExperimentHandler(app).RegisterHandlers()
	handlers.NewUserHandler(app).RegisterHandlers()
	handlers.NewModelHandler(app).RegisterHandlers()
	handlers.NewExportHandler(app).RegisterHandlers()
	app.Listen(":8081")
}
