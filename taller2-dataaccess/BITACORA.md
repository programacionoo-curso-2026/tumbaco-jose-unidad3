\# Bitácora de desarrollo - Taller 2 Data Access



\## Fase 1 - Creación de la entidad



Se creó el paquete `model` y la entidad `Videojuego`.



La entidad contiene los atributos:



\- ID

\- Nombre

\- Genero

\- Plataforma

\- Precio



También se implementaron los métodos `MostrarInformacion()` y `ActualizarPrecio()`.



\### Resultado



La entidad quedó definida y lista para ser utilizada por la capa de acceso a datos.



\---



\## Fase 2 - Conexión con SQLite



Se creó el paquete `dataaccess` y se implementó la conexión con SQLite utilizando `database/sql` y `modernc.org/sqlite`.



También se implementó el método `CrearTabla()` para crear la tabla `videojuegos`.



\### Resultado



La aplicación logró conectarse correctamente a SQLite y crear la tabla.



\---



\## Fase 3 - Primer error en el modelo



Durante la ejecución apareció el error:



```text

model\\videojuego.go:1:1: expected 'package', found 'EOF'

El problema se produjo porque el archivo videojuego.go estaba vacío.



Solución



Se reconstruyó el archivo videojuego.go, comenzando correctamente con:



package model



Después de guardar el archivo se volvió a ejecutar:



go run .



La conexión con SQLite y la creación de la tabla funcionaron correctamente.



Fase 4 - Error en los imports



Al modificar main.go apareció un error de sintaxis relacionado con los imports.



Solución



Se corrigió la sección import y se agregaron correctamente los paquetes:



"taller2-dataaccess/dataaccess"

"taller2-dataaccess/model"



Después de la corrección el programa volvió a compilar.



Fase 5 - Inserción de videojuegos



Se implementó CrearVideojuego() para insertar registros en la tabla videojuegos.



Se realizó una prueba insertando:



Nombre: Minecraft

Género: Sandbox

Plataforma: PC

Precio: 29.99

Resultado



La inserción se realizó correctamente.



Fase 6 - Consulta de videojuegos



Se implementó ObtenerVideojuegoPorID() utilizando una consulta SELECT, QueryRow() y Scan().



También se implementó ObtenerVideojuegos() para recuperar todos los registros.



Resultado



Se pudieron recuperar correctamente los videojuegos almacenados en SQLite.



Fase 7 - Problema con el ID



Durante las pruebas se intentó consultar siempre el videojuego con ID 1.



Después de eliminar ese registro, una nueva ejecución produjo:



sql: no rows in result set

Solución



Se modificó CrearVideojuego() para utilizar:



LastInsertId()



De esta forma, el programa obtiene el ID generado por SQLite y utiliza ese ID para consultar, actualizar y eliminar el videojuego recién creado.



Resultado



El problema quedó solucionado y las operaciones posteriores utilizaron el ID correcto.



Fase 8 - Actualización y eliminación



Se implementaron:



ActualizarPrecio()

EliminarVideojuego()



Se comprobó que el precio del videojuego cambiara de:



29.99



a:



39.99



Posteriormente se eliminó el registro utilizado en la prueba.



Resultado



Las operaciones de actualización y eliminación funcionaron correctamente.



Fase 9 - Formateo y pruebas finales



Se ejecutó:



go fmt ./...



para formatear el código fuente.



Después se ejecutó:



go run .



La aplicación completó correctamente las operaciones de creación, consulta, actualización y eliminación.



Conclusión



El proyecto terminó con una estructura separada entre modelo y acceso a datos.



La aplicación permite gestionar videojuegos mediante SQLite y las operaciones principales de persistencia fueron comprobadas durante el desarrollo.

