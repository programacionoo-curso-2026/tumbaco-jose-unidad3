# Taller 22 - Goroutines

## Nombre del proyecto

`taller22-goroutines`

## Descripción

Este proyecto corresponde al Taller 22 y tiene como propósito aplicar el
concepto de Goroutines en Go.

El programa simula el procesamiento de 20 órdenes y permite observar la
evolución del programa desde una ejecución secuencial hasta una ejecución
concurrente utilizando Goroutines y `sync.WaitGroup`.

## Tecnologías utilizadas

- Go
- Goroutines
- `sync.WaitGroup`
- `fmt`
- `math/rand`
- `time`

---

# Iteración 1 - Generación de órdenes

En esta primera iteración se creó la estructura `Order` y la función
`generateOrders`, encargada de generar 20 órdenes con estado inicial
`pending`.

La ejecución fue correcta.

### Evidencia de ejecución

```text
Todas las operaciones completadas. Finalizando

Iteración 2 - Procesamiento de órdenes

Se agregó la función processOrders, que recorre las 20 órdenes y simula
un tiempo de procesamiento aleatorio mediante time.Sleep.

Evidencia de ejecución
Procesando orden 1
Procesando orden 2
Procesando orden 3
Procesando orden 4
Procesando orden 5
Procesando orden 6
Procesando orden 7
Procesando orden 8
Procesando orden 9
Procesando orden 10
Procesando orden 11
Procesando orden 12
Procesando orden 13
Procesando orden 14
Procesando orden 15
Procesando orden 16
Procesando orden 17
Procesando orden 18
Procesando orden 19
Procesando orden 20
Todas las operaciones completadas. Finalizando
Iteración 3 - Actualización de estados

Se agregó la función updateOrderStatuses.

Cada orden recibe aleatoriamente uno de los siguientes estados:

Procesando
Despachando
Entregado
Evidencia de ejecución
Procesando orden 1
Procesando orden 2
...
Procesando orden 20

Actualizando orden 1 con estado: Despachando
Actualizando orden 2 con estado: Despachando
Actualizando orden 3 con estado: Procesando
Actualizando orden 4 con estado: Entregado
Actualizando orden 5 con estado: Despachando
Actualizando orden 6 con estado: Entregado
Actualizando orden 7 con estado: Entregado
Actualizando orden 8 con estado: Entregado
Actualizando orden 9 con estado: Procesando
Actualizando orden 10 con estado: Entregado
...
Actualizando orden 20 con estado: Despachando

Todas las operaciones completadas. Finalizando

Los estados son aleatorios, por lo que pueden cambiar en cada ejecución.

Iteración 4 - Reporte de estados

Se agregó reportOrderStatus, que muestra el estado de las 20 órdenes
cinco veces, esperando un segundo entre cada reporte.

Evidencia de ejecución
--- Reporte Estado de las Ordenes ---
Orden 1 Entregado
Orden 2 Entregado
Orden 3 Procesando
Orden 4 Procesando
Orden 5 Despachando
Orden 6 Entregado
...
Orden 20 Despachando
---------------------------------------

El reporte anterior se generó cinco veces durante la ejecución.

El programa terminó correctamente:

Todas las operaciones completadas. Finalizando
Iteración 5 - Ejecución secuencial completa

En esta iteración se ejecutaron las tres operaciones principales de forma
secuencial:

processOrders(orders)
updateOrderStatuses(orders)
reportOrderStatus(orders)
Resultado

Primero se procesaron las 20 órdenes, posteriormente se actualizaron sus
estados y finalmente se generaron los cinco reportes.

La ejecución terminó correctamente con:

Todas las operaciones completadas. Finalizando
Iteración 6 - Introducción de Goroutines

En esta iteración se utilizaron Goroutines para ejecutar las tres funciones:

go processOrders(orders)
go updateOrderStatuses(orders)
go reportOrderStatus(orders)
Evidencia de ejecución

La ejecución terminó inmediatamente mostrando:

Todas las operaciones completadas. Saliendo
Observación

Las Goroutines fueron iniciadas, pero la función main terminó sin esperar
a que finalizaran.

Esto demuestra la necesidad de utilizar un mecanismo de sincronización.

Iteración 7 - Introducción de WaitGroup

En esta iteración se agregó:

var wg sync.WaitGroup
wg.Add(3)

y:

wg.Wait()

Las tres funciones se ejecutaron mediante Goroutines.

Evidencia de ejecución

Durante la ejecución se observaron las operaciones de manera concurrente:

Procesando orden 1
Actualizando orden 1 con estado: Entregado
Actualizando orden 2 con estado: Despachando
Procesando orden 2
Procesando orden 3
...
--- Reporte Estado de las Ordenes ---
...

Sin embargo, el programa terminó con:

fatal error: all goroutines are asleep - deadlock!
Explicación

Se utilizó:

wg.Add(3)

pero las Goroutines no ejecutaban:

wg.Done()

Por esta razón el contador del WaitGroup nunca llegó a cero y
wg.Wait() permaneció esperando.

Esta iteración permitió identificar el problema de sincronización.

Iteración 8 - Corrección con WaitGroup

En la última iteración se corrigió el problema mediante funciones anónimas
y defer wg.Done().

Ejemplo:

go func() {
    defer wg.Done()
    processOrders(orders)
}()

La misma estructura se aplicó a las tres Goroutines.

Resultado

Las tres operaciones se ejecutaron concurrentemente y el programa esperó
correctamente hasta que todas terminaron.

Evidencia de ejecución

Durante la ejecución se observaron mensajes intercalados:

Actualizando orden 1 con estado: Entregado
Procesando orden 1
Procesando orden 2
Actualizando orden 2 con estado: Procesando
Actualizando orden 3 con estado: Despachando
Procesando orden 3
Procesando orden 4

--- Reporte Estado de las Ordenes ---
...

Los reportes se ejecutaron mientras las otras Goroutines continuaban
trabajando.

Al finalizar todas las operaciones se obtuvo:

Actualizando orden 17 con estado: Despachando
Actualizando orden 18 con estado: Despachando
Actualizando orden 19 con estado: Entregado
Actualizando orden 20 con estado: Entregado
Todas las operaciones completadas. Saliendo

No se produjo el error de deadlock de la Iteración 7.

Conclusiones

El taller permitió observar progresivamente el funcionamiento de las
Goroutines en Go.

En las primeras iteraciones el programa se ejecutaba de manera secuencial.
En la Iteración 6 se introdujo la ejecución concurrente mediante Goroutines,
pero se observó que main podía finalizar antes que las operaciones.

En la Iteración 7 se incorporó sync.WaitGroup, pero se produjo un deadlock
debido a que no se llamó wg.Done().

Finalmente, en la Iteración 8 se solucionó el problema utilizando
defer wg.Done() dentro de cada Goroutine.

La versión final permite ejecutar las operaciones concurrentemente y esperar
correctamente hasta que todas las Goroutines hayan terminado.
