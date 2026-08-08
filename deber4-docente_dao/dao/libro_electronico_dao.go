package dao

import (
	"database/sql"
)

type LibroElectronicoDAO struct {
	DB *sql.DB
}

func NewLibroElectronicoDAO(db *sql.DB) *LibroElectronicoDAO {
	return &LibroElectronicoDAO{
		DB: db,
	}
}

func (dao *LibroElectronicoDAO) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS libros_electronicos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		autor TEXT NOT NULL,
		genero TEXT NOT NULL,
		isbn TEXT NOT NULL UNIQUE,
		formato TEXT NOT NULL,
		precio REAL NOT NULL
	);
	`

	_, err := dao.DB.Exec(query)
	return err
}
