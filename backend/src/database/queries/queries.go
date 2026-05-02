package queries

var CreateTableReservaciones string = `CREATE TABLE IF NOT EXISTS reservaciones (
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(30) NOT NULL,
	cantidad_personas INT NOT NULL,
	personalizacion VARCHAR(30) NOT NULL,
	CHECK(personalizacion IN ('cita','aniversario','junta de negocios','cumpleaños','salida casual')),
	telefono TEXT NOT NULL,
	fecha TEXT NOT NULL,
	crate_at DATE DEFAULT CURRENT_DATE
)	
`
var CreateTablePlatillos string = `CREATE TABLE IF NOT EXISTS platillos (
	id SERIAL PRIMARY KEY,
	nombre VARCHAR(150) NOT NULL,
	precio FLOAT NOT NULL,
	description VARCHAR(250),
	top INT,
	seccion TEXT NOT NULL,
	CHECK (seccion IN ('huevos','carne','vegetales','lacteos'))
)`

// En desarrollo
var CreateTablePersonalizacion string = `CREATE TABLE personalizacion(
	id SERIAL PRIMARY KEY,
	id_reservacion REFERENCES reservaciones(id) ON DELETE CASCADE,

)
`
