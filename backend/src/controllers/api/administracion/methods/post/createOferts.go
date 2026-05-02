package post

import (
	"App/database/connection"
	"App/models"

	"github.com/gofiber/fiber/v3"
)

func CreateOferts(c fiber.Ctx) error {
	data := new(models.Platillos)
	err := c.Bind().Body(data)
	if err != nil {
		return err
	}
	_, err = connection.DB.Exec(`INSERT INTO platillos (
		nombre,
		precio,
		descripcion,
		top,
		seccion
		) 
		VALUES
		(?,?,?,?,?)`,
		data.Nombre,
		data.Precio,
		data.Descripcion,
		data.Top,
		data.Seccion)
	if err != nil {
		return err
	}
	return c.SendStatus(200)
}
