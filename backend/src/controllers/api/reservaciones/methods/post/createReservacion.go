package post

import (
	"App/database/connection"
	"App/models"

	"github.com/gofiber/fiber/v3"
)

func CreateReservacion(c fiber.Ctx) error {
	data := new(models.Reservaciones)
	err := c.Bind().Body(data)
	if err != nil {
		return err
	}
	_, err = connection.DB.Exec(`INSERT INTO reservaciones (
		nombre,
		cantidad_personas,
		personalizacion,
		telefono,
		fecha
		) 
		VALUES
		(?,?,?,?,?)`,
		data.Nombre,
		data.Cantidad_personas,
		data.Personalizacion,
		data.Telefono,
		data.Fecha)
	if err != nil {
		return err
	}
	return c.SendStatus(200)
}
