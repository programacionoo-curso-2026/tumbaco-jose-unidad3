# Bitácora de avances — deber4-docente_dao

## 1. Inicio del proyecto

Se creó la carpeta `deber4-docente_dao` para desarrollar la actividad de acceso a datos utilizando Go, SQLite y el patrón DAO.

Se organizó el proyecto utilizando los paquetes:

* `model`: contiene la entidad `LibroElectronico`.
* `dataaccess`: administra la conexión con SQLite.
* `dao`: contiene las operaciones de acceso a datos.
* `main.go`: ejecuta y demuestra las operaciones del sistema.

## 2. Creación del modelo

Se desarrolló el modelo `LibroElectronico` con los siguientes atributos:

* `ID`
* `Titulo`
* `Autor`
* `Genero`
* `ISBN`
* `Formato`
* `Precio`

También se implementaron métodos relacionados con la información y actualización del libro.

## 3. Configuración de la base de datos

Se configuró la conexión con SQLite mediante el paquete `dataaccess`.

La aplicación utiliza la base de datos local:

`libros_electronicos.db`

Se implementó la inicialización de la base de datos y posteriormente se verificó que la conexión funcionara correctamente.

## 4. Implementación del DAO

Se creó `LibroElectronicoDAO` dentro del paquete `dao`.

El DAO permite realizar las siguientes operaciones:

* Crear la tabla.
* Crear o insertar libros.
* Consultar un libro por ID.
* Consultar todos los libros.
* Actualizar un libro.
* Eliminar un libro.

La separación entre modelo, conexión y DAO permite organizar mejor las responsabilidades del programa.

## 5. Implementación del programa principal

Se configuró `main.go` para:

1. Inicializar la base de datos.
2. Crear el DAO.
3. Crear la tabla.
4. Insertar un libro.
5. Buscar el libro por ID.
6. Obtener todos los libros.
7. Actualizar el precio.
8. Comprobar la actualización.
9. Eliminar el libro.

Como registro de prueba se utilizó el libro:

* Título: `Clean Code`
* Autor: `Robert C. Martin`
* Género: `Programación`
* ISBN: `9780132350884`
* Formato: `PDF`
* Precio inicial: `29.99`

## 6. Prueba de funcionamiento

Se ejecutó:

```text
go run .
```

La ejecución fue exitosa y mostró:

```text
Base de datos inicializada correctamente
Tabla de libros electrónicos creada correctamente
Libro creado correctamente con ID: 1
Libro encontrado: &{ID:1 Titulo:Clean Code Autor:Robert C. Martin Genero:Programación ISBN:9780132350884 Formato:PDF Precio:29.99}
Lista de libros:
ID: 1 | Título: Clean Code | Autor: Robert C. Martin | Género: Programación | ISBN: 9780132350884 | Formato: PDF | Precio: 29.99
Libro actualizado correctamente
Libro actualizado: &{ID:1 Titulo:Clean Code Autor:Robert C. Martin Genero:Programación ISBN:9780132350884 Formato:PDF Precio:39.99}
Libro eliminado correctamente
```

La prueba permitió comprobar que las operaciones principales del DAO funcionan correctamente.

## 7. Verificación de las operaciones CRUD

Se verificaron las siguientes operaciones:

* **Create:** se creó correctamente un libro con ID `1`.
* **Read:** se consultó correctamente el libro por su ID.
* **Read All:** se obtuvo correctamente la lista de libros.
* **Update:** el precio se actualizó de `29.99` a `39.99`.
* **Delete:** el libro fue eliminado correctamente.

## 8. Advertencia durante Git

Al realizar el commit de `main.go`, Git mostró la siguiente advertencia:

```text
warning: in the working copy of 'deber4-docente_dao/main.go',
LF will be replaced by CRLF the next time Git touches it
```

Esta advertencia está relacionada con los finales de línea del archivo en el entorno de Windows. No impidió realizar el commit ni afectó la ejecución del programa.

El commit se realizó correctamente con el mensaje:

```text
deber4: demuestra operaciones CRUD en main
```

## 9. Documentación del proyecto

Se creó el archivo `README.md` para documentar:

* Objetivo del proyecto.
* Tecnologías utilizadas.
* Estructura del proyecto.
* Modelo de datos.
* Operaciones del DAO.
* Base de datos.
* Instrucciones de ejecución.
* Evidencias de funcionamiento.
* Decisiones de diseño.

Posteriormente, el README fue agregado al repositorio mediante Git.

El commit correspondiente fue:

```text
deber4: documenta ejecucion y decisiones del proyecto
```

## 10. Estado actual

El proyecto cuenta con:

* Modelo de libros electrónicos.
* Conexión con SQLite.
* DAO funcional.
* Operaciones CRUD comprobadas.
* `main.go` funcional.
* `README.md` documentado.
* `BITACORA.md` con el registro del proceso.

Hasta este punto, las pruebas realizadas indican que el programa funciona correctamente y que el repositorio mantiene un historial incremental de cambios.
