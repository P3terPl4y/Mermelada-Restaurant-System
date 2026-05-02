package models

type Reservaciones struct {
	Id                int    `json:"id"`
	Nombre            string `json:"nombre"`
	Cantidad_personas int    `json:"cantidad_personas"`
	Personalizacion   string `json:"personalizacion"`
	Telefono          string `json:"telefono"`
	Fecha             string `json:"fecha"`
	Create_at         string `json:"create_at"`
}
type Platillos struct {
	Id          int     `json:"id"`
	Nombre      string  `json:"nombre"`
	Precio      float64 `json:"precio"`
	Descripcion string  `json:"descripcion"`
	Top         int     `json:"top"`
	Seccion     string  `json:"seccion"`
}
