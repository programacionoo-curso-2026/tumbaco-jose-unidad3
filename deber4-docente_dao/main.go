package main

import (
	"deber4-docente_dao/dao"
	"deber4-docente_dao/dataaccess"
	"deber4-docente_dao/model"
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

	// Crear un libro
	libroID, err := libroDAO.Create(
		model.LibroElectronico{
			Titulo:  "Clean Code",
			Autor:   "Robert C. Martin",
			Genero:  "Programación",
			ISBN:    "9780132350884",
			Formato: "PDF",
			Precio:  29.99,
		},
	)
	if err != nil {
		log.Fatalf("Error al crear libro: %v", err)
	}

	log.Printf("Libro creado correctamente con ID: %d", libroID)

	// Obtener un libro por ID
	libro, err := libroDAO.GetByID(libroID)
	if err != nil {
		log.Fatalf("Error al buscar libro: %v", err)
	}

	log.Printf("Libro encontrado: %+v", libro)

	// Obtener todos los libros
	libros, err := libroDAO.GetAll()
	if err != nil {
		log.Fatalf("Error al obtener libros: %v", err)
	}

	log.Println("Lista de libros:")

	for _, libro := range libros {
		log.Printf(
			"ID: %d | Título: %s | Autor: %s | Género: %s | ISBN: %s | Formato: %s | Precio: %.2f",
			libro.ID,
			libro.Titulo,
			libro.Autor,
			libro.Genero,
			libro.ISBN,
			libro.Formato,
			libro.Precio,
		)
	}

	// Actualizar el libro
	libroActualizado := *libro
	libroActualizado.Precio = 39.99

	if err := libroDAO.Update(libroActualizado); err != nil {
		log.Fatalf("Error al actualizar libro: %v", err)
	}

	log.Println("Libro actualizado correctamente")

	// Comprobar actualización
	libro, err = libroDAO.GetByID(libroID)
	if err != nil {
		log.Fatalf("Error al comprobar actualización: %v", err)
	}

	log.Printf("Libro actualizado: %+v", libro)

	// Eliminar el libro
	if err := libroDAO.Delete(libroID); err != nil {
		log.Fatalf("Error al eliminar libro: %v", err)
	}

	log.Println("Libro eliminado correctamente")
}
