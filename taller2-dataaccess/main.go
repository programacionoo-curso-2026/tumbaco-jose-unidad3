package main

import (
	"fmt"
	"log"

	"taller2-dataaccess/dataaccess"
	"taller2-dataaccess/model"
)

func main() {
	da := dataaccess.NewDataAccess()
	defer da.Close()

	if da.DB == nil {
		log.Fatal("La conexión no fue creada")
	}

	fmt.Println("Conexión a SQLite realizada correctamente.")

	err := da.CrearTabla()
	if err != nil {
		log.Fatal("Error al crear la tabla:", err)
	}

	fmt.Println("Tabla videojuegos creada correctamente.")

	videojuego := model.Videojuego{
		Nombre:     "Minecraft",
		Genero:     "Sandbox",
		Plataforma: "PC",
		Precio:     29.99,
	}

	idCreado, err := da.CrearVideojuego(videojuego)
	if err != nil {
		log.Fatal("Error al crear el videojuego:", err)
	}

	fmt.Printf("Videojuego creado correctamente con ID: %d\n", idCreado)

	videojuegoEncontrado, err := da.ObtenerVideojuegoPorID(idCreado)
	if err != nil {
		log.Fatal("Error al buscar el videojuego:", err)
	}

	fmt.Printf("Videojuego encontrado: %+v\n", videojuegoEncontrado)

	videojuegos, err := da.ObtenerVideojuegos()
	if err != nil {
		log.Fatal("Error al obtener los videojuegos:", err)
	}

	fmt.Println("\nLista de videojuegos:")

	for _, v := range videojuegos {
		fmt.Printf(
			"ID: %d | Nombre: %s | Género: %s | Plataforma: %s | Precio: %.2f\n",
			v.ID,
			v.Nombre,
			v.Genero,
			v.Plataforma,
			v.Precio,
		)
	}

	err = da.ActualizarPrecio(idCreado, 39.99)
	if err != nil {
		log.Fatal("Error al actualizar el precio:", err)
	}

	fmt.Println("\nPrecio actualizado correctamente.")

	videojuegoActualizado, err := da.ObtenerVideojuegoPorID(idCreado)
	if err != nil {
		log.Fatal("Error al consultar el videojuego actualizado:", err)
	}

	fmt.Printf("Videojuego actualizado: %+v\n", videojuegoActualizado)

	err = da.EliminarVideojuego(idCreado)
	if err != nil {
		log.Fatal("Error al eliminar el videojuego:", err)
	}

	fmt.Println("Videojuego eliminado correctamente.")
}
