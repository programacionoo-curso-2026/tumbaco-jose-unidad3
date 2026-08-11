# Taller 23 - Goroutines y Mutex

## Datos del taller

- **Directorio:** `taller23-goroutines`
- **Lenguaje:** Go
- **Tema:** Goroutines, concurrencia y Mutex
- **Módulo:** `taller23-goroutines`

## Objetivo

Aplicar el uso de Goroutines y mecanismos de sincronización mediante `sync.WaitGroup` y `sync.Mutex`, trabajando con múltiples actualizaciones concurrentes sobre órdenes.

El programa administra 20 órdenes y ejecuta tres Goroutines que actualizan el estado de cada orden.

---

# Iteración 1

En la primera iteración se implementó la estructura `Order` incorporando un `sync.Mutex` propio para controlar el acceso concurrente al estado de cada orden.

También se declaró un contador global de actualizaciones protegido mediante otro `sync.Mutex`.

### Estructura utilizada

```go
type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)
Evidencia de ejecución

Comando utilizado:

go fmt ./...
go test ./...
go run .

Resultado:

Taller 23 - Goroutines y Mutex
Iteración 1

La primera iteración compiló y ejecutó correctamente.

Iteración 2

En esta iteración se incorporaron tres Goroutines para actualizar concurrentemente las 20 órdenes.

Se utilizó sync.WaitGroup para esperar la finalización de las tres Goroutines.

Funcionamiento

Cada Goroutine recorre las 20 órdenes y ejecuta:

updateOrderStatus(order)

El programa permite observar que las actualizaciones se realizan de manera concurrente.

Evidencia de ejecución

Comandos utilizados:

go fmt ./...
go test ./...
go run .

Resultado representativo:

Actualizando orden 1 con estado: Despachando
Actualizando orden 1 con estado: Entregado
Actualizando orden 1 con estado: Entregado
Actualizando orden 2 con estado: Entregado
Actualizando orden 2 con estado: Procesando
Actualizando orden 3 con estado: Despachando
Actualizando orden 3 con estado: Entregado
...
Actualizando orden 20 con estado: Despachando
Actualizando orden 20 con estado: Entregado
Actualizando orden 20 con estado: Despachando
Todas las operaciones completadas. Saliendo
Total Actualizaciones 0

En esta etapa el contador todavía mostraba 0, debido a que el incremento de totalUpdates todavía no había sido implementado.

Iteración 3

En la tercera iteración se agregó el control de acceso mediante Mutex para el contador global totalUpdates.

Se implementó:

updateMutex.Lock()
defer updateMutex.Unlock()

currentUpdates := totalUpdates
time.Sleep(5 * time.Millisecond)
totalUpdates = currentUpdates + 1

También se mantuvo el Mutex propio de cada orden para proteger la modificación de su estado.

Resultado esperado

Existen:

20 órdenes.
3 Goroutines.
1 actualización por orden en cada Goroutine.

Por lo tanto:

20 × 3 = 60 actualizaciones
Evidencia de ejecución

Comandos utilizados:

go fmt ./...
go test ./...
go run .

Resultado final:

Actualizando orden 1 con estado: Procesando
Actualizando orden 1 con estado: Entregado
Actualizando orden 2 con estado: Procesando
...
Actualizando orden 20 con estado: Despachando
Actualizando orden 20 con estado: Procesando
Todas las operaciones completadas. Saliendo
Total Actualizaciones 60

La ejecución terminó correctamente y el contador registró las 60 actualizaciones esperadas.

Conceptos aplicados
Goroutines

Se utilizaron tres Goroutines para realizar simultáneamente el procesamiento de las órdenes:

for i := 0; i < 3; i++ {
	go func() {
		defer wg.Done()

		for _, order := range orders {
			updateOrderStatus(order)
		}
	}()
}
WaitGroup

Se utilizó sync.WaitGroup para esperar que las tres Goroutines terminaran:

var wg sync.WaitGroup
wg.Add(3)

Cada Goroutine utiliza:

defer wg.Done()

Y el programa espera mediante:

wg.Wait()
Mutex

Cada orden dispone de su propio Mutex:

mu sync.Mutex

Esto permite controlar el acceso al estado de la orden.

También se utiliza otro Mutex para proteger el contador global:

updateMutex sync.Mutex

De esta manera se evita que varias Goroutines modifiquen simultáneamente el valor de totalUpdates.

Validaciones realizadas

Se ejecutaron las siguientes instrucciones:

go fmt ./...

Resultado:

El código fue formateado correctamente.

También se ejecutó:

go test ./...

Resultado:

?       taller23-goroutines     [no test files]

Finalmente se ejecutó:

go run .

El programa terminó correctamente con:

Todas las operaciones completadas. Saliendo
Total Actualizaciones 60
Conclusión

El taller permitió aplicar concurrencia mediante Goroutines y sincronización mediante sync.WaitGroup y sync.Mutex.

La ejecución final demostró que las tres Goroutines realizaron correctamente las actualizaciones de las 20 órdenes, obteniendo un total de:

60 actualizaciones

El programa finalizó correctamente después de que todas las Goroutines completaron su trabajo.
