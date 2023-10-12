package main

import (
	"github.com/elwicho/APIREST-GO/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Llama a la función ConnectDb() para establecer la conexión a la base de datos.
	database.ConnectDb()

	// Crea una nueva instancia de la aplicación Fiber.
	app := fiber.New()

	// Configura las rutas de la aplicación llamando a la función setupRoutes.
	setupRoutes(app)

	// Inicia la aplicación y la hace escuchar en el puerto 8000.
	app.Listen(":8000")
}
