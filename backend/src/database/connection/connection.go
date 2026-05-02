package connection

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	var err error
	//Version de desarrollo con sqlite3
	DB, err = sql.Open("sqlite3", "database/mermelada.db")
	if err != nil {
		return err
	}

	var CreateTableReservaciones string = `CREATE TABLE IF NOT EXISTS reservaciones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nombre VARCHAR(30) NOT NULL,
    cantidad_personas INT NOT NULL,
    personalizacion VARCHAR(30) NOT NULL,
    telefono TEXT NOT NULL,
    fecha TEXT NOT NULL,
	create_at TEXT DEFAULT (CURRENT_DATE)
);`
	var CreateTablePlatillos string = `CREATE TABLE IF NOT EXISTS platillos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nombre VARCHAR(150) NOT NULL,
	precio FLOAT NOT NULL,
	descripcion VARCHAR(250),
	top INT,
	seccion TEXT NOT NULL
)`
	trans, err := DB.Begin()
	_, err = trans.Exec(CreateTableReservaciones)
	if err != nil {
		trans.Rollback()
		return err
	}
	_, err = trans.Exec(CreateTablePlatillos)
	if err != nil {
		trans.Rollback()
		return err
	}
	err = trans.Commit()
	if err != nil {
		trans.Rollback()
		return err
	}
	//Para PostgreSQL
	/*	trans, err := DB.Begin()
		_, err = trans.Exec(queries.CreateTableReservaciones)
		if err != nil {
			trans.Rollback()
			return err
		}
		err = trans.Commit()
		if err != nil {
			trans.Rollback()
			return err
		}
	*/
	return nil
}
