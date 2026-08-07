package main

import (
	"fmt"
	"log"

	"taller2-dataaccess/dataaccess"
)

func main() {
	da := dataaccess.NewDataAccess()
	defer da.Close()

	fmt.Println("Conexión a SQLite realizada correctamente.")

	if da.DB == nil {
		log.Fatal("La conexión no fue creada")
	}
}
