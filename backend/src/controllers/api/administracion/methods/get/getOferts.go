package get

import (
	"App/database/connection"
	"App/models"

	"github.com/gofiber/fiber/v3"
)

func GetOferts(c fiber.Ctx) error {
	section := c.Params("section")
	rows, err := connection.DB.Query("SELECT * FROM platillos WHERE seccion=?", section)
	if err != nil {
		return err
	}
	defer rows.Close()
	var platillos_list []models.Platillos
	for rows.Next() {
		var row models.Platillos
		err := rows.Scan(&row.Id, &row.Nombre, &row.Precio, &row.Descripcion, &row.Top, &row.Seccion)
		if err != nil {
			return err
		}
		platillos_list = append(platillos_list, row)
	}
	return c.JSON(platillos_list)
}
