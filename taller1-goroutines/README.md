# Taller 1 - Goroutines

## Objetivo

Comprender el funcionamiento básico de las goroutines en Go y observar
cómo se ejecutan varias tareas de manera concurrente.

## Descripción

En este taller se trabaja progresivamente con goroutines:

1. Ejecución normal de una función.
2. Ejecución de una función como goroutine utilizando `go`.
3. Uso de `time.Sleep` para permitir que la goroutine termine.
4. Ejecución de múltiples goroutines mediante un ciclo `for`.
5. Uso de retrasos aleatorios con `math/rand`.

## Ejecución

Desde la raíz del repositorio se puede ejecutar el programa con:

```bash
go run ./taller1-goroutines/src/ejemplo1.go

