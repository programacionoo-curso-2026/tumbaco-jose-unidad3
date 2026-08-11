package main

import (
	"deber5-docente_dao/dao"
	"deber5-docente_dao/dataaccess"
	"deber5-docente_dao/model"
	"log"
)

func main() {
	// Inicializar la base de datos
	db := dataaccess.InitDB()
	defer db.Close()

	log.Println("Base de datos inicializada correctamente")

	// Crear el DAO
	libroDAO := dao.NewLibroElectronicoDAO(db)

	// 1. Crear la tabla
	if err := libroDAO.CreateTable(); err != nil {
		log.Fatalf("Error al crear tabla: %v", err)
	}

	log.Println("Tabla de libros electrónicos creada correctamente")

	// 2. INSERTAR un libro
	libro := &model.LibroElectronico{
		Titulo:  "Clean Code",
		Autor:   "Robert C. Martin",
		Genero:  "Programación",
		ISBN:    "9780132350884",
		Formato: "PDF",
		Precio:  29.99,
	}

	if err := libroDAO.Create(libro); err != nil {
		log.Fatalf("Error al crear libro: %v", err)
	}

	log.Printf("Libro creado correctamente con ID: %d", libro.ID)

	// 3. BUSCAR un libro por ID
	libroEncontrado, err := libroDAO.GetByID(libro.ID)
	if err != nil {
		log.Fatalf("Error al buscar libro: %v", err)
	}

	log.Printf("Libro encontrado: %+v", libroEncontrado)

	// 4. OBTENER TODOS los libros
	libros, err := libroDAO.GetAll()
	if err != nil {
		log.Fatalf("Error al obtener libros: %v", err)
	}

	log.Println("Lista de libros:")

	for _, l := range libros {
		log.Printf(
			"ID: %d | Título: %s | Autor: %s | Género: %s | ISBN: %s | Formato: %s | Precio: %.2f",
			l.ID,
			l.Titulo,
			l.Autor,
			l.Genero,
			l.ISBN,
			l.Formato,
			l.Precio,
		)
	}

	// 5. ACTUALIZAR el precio
	libro.Precio = 39.99

	if err := libroDAO.Update(libro); err != nil {
		log.Fatalf("Error al actualizar libro: %v", err)
	}

	log.Println("Libro actualizado correctamente")

	// Comprobar la actualización
	libroActualizado, err := libroDAO.GetByID(libro.ID)
	if err != nil {
		log.Fatalf("Error al comprobar actualización: %v", err)
	}

	log.Printf("Libro actualizado: %+v", libroActualizado)

	// 6. ELIMINAR el libro
	if err := libroDAO.Delete(libro.ID); err != nil {
		log.Fatalf("Error al eliminar libro: %v", err)
	}

	log.Println("Libro eliminado correctamente")
}
