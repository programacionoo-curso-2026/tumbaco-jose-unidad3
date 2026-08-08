package main

import (
	"deber4-docente_dao/dao"
	"deber4-docente_dao/dataaccess"
	"log"
)

func main() {
	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close()

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	libroDAO := dao.NewLibroElectronicoDAO(db)

	// Crear la tabla
	if err := libroDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}

	log.Println("Tabla de libros electrónicos creada correctamente")
}
