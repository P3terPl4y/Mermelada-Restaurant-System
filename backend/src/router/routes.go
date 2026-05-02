package router

import (
	administracion_get "App/controllers/api/administracion/methods/get"
	administracion_post "App/controllers/api/administracion/methods/post"
	reservaciones_get "App/controllers/api/reservaciones/methods/get"
	reservaciones_post "App/controllers/api/reservaciones/methods/post"

	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {

	app.Get("/api/admin/info", reservaciones_get.GetInfo)
	app.Get("/api/oferts/:section", administracion_get.GetOferts)
	app.Post("/api/user/reservar", reservaciones_post.CreateReservacion)
	app.Post("/api/admin/create/platillo", administracion_post.CreateOferts)

}
