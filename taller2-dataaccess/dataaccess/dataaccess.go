package dataaccess

import (
	"database/sql"
	"log"

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

func (da *DataAccess) Close() {
	da.DB.Close()
}
