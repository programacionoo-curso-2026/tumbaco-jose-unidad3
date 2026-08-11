# Bitácora de Desarrollo - Taller 23 Goroutines

## Información general

- **Proyecto:** Taller 23 - Goroutines y Mutex
- **Directorio:** `taller23-goroutines`
- **Lenguaje:** Go
- **Tema:** Concurrencia, Goroutines y Mutex

---

## Iteración 1 - Estructura inicial

### Actividad realizada

Se creó la estructura `Order` para representar las órdenes del sistema.

Se agregó un `sync.Mutex` dentro de cada orden para controlar el acceso concurrente a su estado.

También se declararon las variables globales:

- `totalUpdates`
- `updateMutex`

### Validación

Se ejecutaron:

```powershell
go fmt ./...
go test ./...
go run .
Resultado

La ejecución fue correcta:

Taller 23 - Goroutines y Mutex
Iteración 1
Iteración 2 - Uso de Goroutines
Actividad realizada

Se implementaron tres Goroutines para procesar las 20 órdenes de manera concurrente.

Se utilizó sync.WaitGroup para controlar la finalización de las tres Goroutines.

Cada Goroutine recorre las 20 órdenes y ejecuta la función:

updateOrderStatus(order)
Observación

La salida mostró que las actualizaciones no necesariamente aparecen en orden secuencial.

Esto permite evidenciar el comportamiento concurrente de las Goroutines.

Validación

Se ejecutaron:

go fmt ./...
go test ./...
go run .
Resultado

El programa terminó correctamente:

Todas las operaciones completadas. Saliendo
Total Actualizaciones 0

El valor 0 era esperado en esta etapa porque todavía no se había implementado el incremento del contador global.

Iteración 3 - Protección del contador
Actividad realizada

Se agregó el bloqueo mediante updateMutex para proteger la variable global totalUpdates.

Se implementó:

updateMutex.Lock()
defer updateMutex.Unlock()

currentUpdates := totalUpdates
time.Sleep(5 * time.Millisecond)
totalUpdates = currentUpdates + 1

De esta forma, una sola Goroutine puede modificar el contador a la vez.

Validación

Se ejecutaron:

go fmt ./...
go test ./...
go run .
Resultado

La ejecución finalizó correctamente con:

Todas las operaciones completadas. Saliendo
Total Actualizaciones 60
Verificación

El resultado de 60 actualizaciones corresponde a:

20 órdenes × 3 Goroutines = 60 actualizaciones
Problemas encontrados

Durante el desarrollo no se presentaron errores de compilación en la implementación final.

En las primeras iteraciones el contador mostraba:

Total Actualizaciones 0

Esto se debía a que todavía no se había agregado la sección encargada de incrementar totalUpdates.

Después de implementar el updateMutex, el contador terminó correctamente en:

Total Actualizaciones 60
Comandos utilizados

Durante el desarrollo se utilizaron principalmente:

go mod init taller23-goroutines
go fmt ./...
go test ./...
go run .

También se utilizó Visual Studio Code para editar los archivos del proyecto.

Resultado final

El taller terminó funcionando correctamente.

Se aplicaron:

Goroutines.
sync.WaitGroup.
sync.Mutex.
Exclusión mutua para modificar estados.
Exclusión mutua para proteger un contador compartido.
Ejecución concurrente de múltiples tareas.

El resultado final fue:

Todas las operaciones completadas. Saliendo
Total Actualizaciones 60
