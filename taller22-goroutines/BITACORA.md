# Bitácora - Taller 22 Goroutines

## Proyecto

`taller22-goroutines`

## Objetivo

Aplicar Goroutines y mecanismos de sincronización en Go mediante el
procesamiento concurrente de órdenes.

---

## Iteración 1

Se creó la estructura `Order` y la función `generateOrders`.

Se generaron 20 órdenes con estado inicial `pending`.

Resultado:

```text
Todas las operaciones completadas. Finalizando

Iteración 2

Se implementó processOrders.

Se utilizó time.Sleep junto con rand.Intn para simular tiempos
aleatorios de procesamiento.

Se comprobó que las 20 órdenes fueran procesadas correctamente.

Resultado:

Procesando orden 1
...
Procesando orden 20
Todas las operaciones completadas. Finalizando
Iteración 3

Se implementó updateOrderStatuses.

Las órdenes recibieron aleatoriamente los estados:

Procesando
Despachando
Entregado

La ejecución fue correcta y se actualizaron las 20 órdenes.

Iteración 4

Se implementó reportOrderStatus.

El programa generó cinco reportes con el estado de las 20 órdenes,
separados por un segundo.

La ejecución finalizó correctamente.

Iteración 5

Se integraron las funciones:

processOrders
updateOrderStatuses
reportOrderStatus

de forma secuencial.

Se verificó el flujo completo del programa.

Iteración 6

Se agregaron Goroutines:

go processOrders(orders)
go updateOrderStatuses(orders)
go reportOrderStatus(orders)

Se comprobó que main terminaba antes de que las Goroutines pudieran
completar sus operaciones.

La salida observada fue:

Todas las operaciones completadas. Saliendo

Esto evidenció la necesidad de sincronización.

Iteración 7

Se agregó:

var wg sync.WaitGroup
wg.Add(3)

y:

wg.Wait()

Las tres funciones comenzaron a ejecutarse concurrentemente.

Sin embargo, no se agregó wg.Done().

Resultado:

fatal error: all goroutines are asleep - deadlock!
Problema identificado

El WaitGroup esperaba que su contador llegara a cero, pero ninguna
Goroutine reducía el contador.

Iteración 8

Se corrigió el problema mediante funciones anónimas:

go func() {
    defer wg.Done()
    processOrders(orders)
}()

Se aplicó el mismo mecanismo a las tres Goroutines.

Resultado final

Las operaciones se ejecutaron concurrentemente y wg.Wait() esperó
correctamente hasta que todas las Goroutines terminaron.

La ejecución terminó con:

Todas las operaciones completadas. Saliendo

No se presentó el deadlock de la Iteración 7.
