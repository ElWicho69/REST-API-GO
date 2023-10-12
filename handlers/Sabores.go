package handlers

import (
	"github.com/elwicho/APIREST-GO/database"
	"github.com/elwicho/APIREST-GO/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Primera Rest API en GO del Wicho :D")
}

// ListarSabores devuelve una lista de sabores (GET /sabores).
func ListarSabores(c *fiber.Ctx) error {
	sabor := []models.Sabores{}
	database.DB.Db.Find(&sabor)

	return c.Status(200).JSON(sabor)
}

// GetSabor obtiene un sabor por su ID (GET /sabores/:id).
func GetSabor(c *fiber.Ctx) error {
	id := c.Params("id")

	// Busca el sabor por ID en la base de datos
	sabor := new(models.Sabores)
	if err := database.DB.Db.First(&sabor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sabor no encontrado",
		})
	}

	return c.Status(fiber.StatusOK).JSON(sabor)
}

// Create crea un nuevo sabor (POST /sabores).
func Create(c *fiber.Ctx) error {
	sabor := new(models.Sabores)
	if err := c.BodyParser(sabor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&sabor)

	return c.Status(200).JSON(sabor)
}

// DeleteSabor elimina un sabor por su ID (DELETE /sabores/:id).
func DeleteSabor(c *fiber.Ctx) error {
	id := c.Params("id")

	// Verifica si el recurso existe antes de eliminarlo
	sabor := new(models.Sabores)
	if err := database.DB.Db.First(&sabor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sabor no encontrado",
		})
	}

	// Elimina el recurso
	if err := database.DB.Db.Delete(&sabor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sabor eliminado correctamente",
	})
}

// PutSabor actualiza un sabor por su ID (PUT /sabores/:id).
func PutSabor(c *fiber.Ctx) error {
	id := c.Params("id")

	// Verifica si el recurso existe antes de actualizarlo
	sabor := new(models.Sabores)
	if err := database.DB.Db.First(&sabor, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Sabor no encontrado",
		})
	}

	// Parsea el cuerpo de la solicitud en un nuevo objeto de sabor
	newSabor := new(models.Sabores)
	if err := c.BodyParser(newSabor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Actualiza el recurso con los datos nuevos
	sabor.Sabor = newSabor.Sabor

	if err := database.DB.Db.Save(&sabor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sabor)
}
