package dataaccess

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "libros_electronicos.db")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	return db
}
