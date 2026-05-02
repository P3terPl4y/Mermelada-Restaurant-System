package get

import (
	"App/database/connection"
	"App/models"

	"github.com/gofiber/fiber/v3"
)

func GetInfo(c fiber.Ctx) error {
	rows, err := connection.DB.Query("SELECT * FROM reservaciones")
	if err != nil {
		return err
	}
	defer rows.Close()
	var reservaciones_list []models.Reservaciones
	for rows.Next() {
		var row models.Reservaciones
		err := rows.Scan(&row.Id, &row.Nombre, &row.Cantidad_personas, &row.Personalizacion, &row.Telefono, &row.Fecha, &row.Create_at)
		if err != nil {
			return err
		}
		reservaciones_list = append(reservaciones_list, row)
	}
	return c.JSON(reservaciones_list)
}
