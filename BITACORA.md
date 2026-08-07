// Punto de entrada del taller
// Reemplaza este archivo con tu código (Go, Java, Python, etc. según indique el docente)

package main

import "fmt"

func main() {
	fmt.Println("Taller de POO — [tu nombre]")
}

## 29/07/2026 - Taller 1

# Bitácora de avances

## Taller 1 - Goroutines

### Fase 1 - Estructura inicial

Se creó la estructura correspondiente al Taller 1 dentro del repositorio.

Se creó la carpeta `taller1-goroutines` con su archivo `README.md` y la
carpeta `src`, donde se encuentra el archivo `ejemplo1.go`.

---

### Fase 2 - Primera implementación

Se implementó la función `ShowGoroutine`, recibiendo un identificador
como parámetro y mostrando el número de la goroutine mediante `fmt.Printf`.

Inicialmente la función fue ejecutada de manera normal desde `main`.

---

### Fase 3 - Uso de goroutines

Se modificó la ejecución de `ShowGoroutine` utilizando la palabra clave
`go`.

Esto permitió ejecutar la función como una goroutine.

Durante esta etapa se observó que el programa podía finalizar antes de
que la goroutine alcanzara a mostrar su resultado.

---

### Fase 4 - Problema y solución

**Problema:**

Al ejecutar:

```go
go ShowGoroutine(1)