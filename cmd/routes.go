package main

import (
	"github.com/elwicho/APIREST-GO/handlers"
	"github.com/gofiber/fiber/v2"
)

// La función setupRoutes configura las rutas para la aplicación Fiber.
func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.Home)

	// Define rutas para la entidad "Sabores" utilizando los controladores proporcionados en el paquete "handlers".
	app.Get("/sabores", handlers.ListarSabores)
	app.Get("/sabores/:id", handlers.GetSabor)
	app.Post("/sabores", handlers.Create)
	app.Delete("/sabores/:id", handlers.DeleteSabor)
	app.Put("/sabores/:id", handlers.PutSabor)
}
