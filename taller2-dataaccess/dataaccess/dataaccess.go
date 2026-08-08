package dataaccess

import (
	"database/sql"
	"log"

	"taller2-dataaccess/model"

	_ "modernc.org/sqlite"
)

type DataAccess struct {
	DB *sql.DB
}

func NewDataAccess() *DataAccess {
	db, err := sql.Open("sqlite", "videojuegos.db")
	if err != nil {
		log.Fatal(err)
	}

	return &DataAccess{
		DB: db,
	}
}

func (da *DataAccess) CrearTabla() error {
	query := `
	CREATE TABLE IF NOT EXISTS videojuegos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		genero TEXT NOT NULL,
		plataforma TEXT NOT NULL,
		precio REAL NOT NULL
	);
	`

	_, err := da.DB.Exec(query)
	return err
}

func (da *DataAccess) CrearVideojuego(v model.Videojuego) (int, error) {
	query := `
	INSERT INTO videojuegos (nombre, genero, plataforma, precio)
	VALUES (?, ?, ?, ?);
	`

	resultado, err := da.DB.Exec(
		query,
		v.Nombre,
		v.Genero,
		v.Plataforma,
		v.Precio,
	)

	if err != nil {
		return 0, err
	}

	id, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (da *DataAccess) ObtenerVideojuegoPorID(id int) (*model.Videojuego, error) {
	query := `
	SELECT id, nombre, genero, plataforma, precio
	FROM videojuegos
	WHERE id = ?;
	`

	var videojuego model.Videojuego

	err := da.DB.QueryRow(query, id).Scan(
		&videojuego.ID,
		&videojuego.Nombre,
		&videojuego.Genero,
		&videojuego.Plataforma,
		&videojuego.Precio,
	)

	if err != nil {
		return nil, err
	}

	return &videojuego, nil
}

func (da *DataAccess) ObtenerVideojuegos() ([]model.Videojuego, error) {
	query := `
	SELECT id, nombre, genero, plataforma, precio
	FROM videojuegos;
	`

	rows, err := da.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videojuegos []model.Videojuego

	for rows.Next() {
		var videojuego model.Videojuego

		err := rows.Scan(
			&videojuego.ID,
			&videojuego.Nombre,
			&videojuego.Genero,
			&videojuego.Plataforma,
			&videojuego.Precio,
		)

		if err != nil {
			return nil, err
		}

		videojuegos = append(videojuegos, videojuego)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videojuegos, nil
}

func (da *DataAccess) ActualizarPrecio(id int, nuevoPrecio float64) error {
	query := `
	UPDATE videojuegos
	SET precio = ?
	WHERE id = ?;
	`

	_, err := da.DB.Exec(query, nuevoPrecio, id)
	return err
}

func (da *DataAccess) EliminarVideojuego(id int) error {
	query := `
	DELETE FROM videojuegos
	WHERE id = ?;
	`

	_, err := da.DB.Exec(query, id)
	return err
}

func (da *DataAccess) Close() {
	da.DB.Close()
}
