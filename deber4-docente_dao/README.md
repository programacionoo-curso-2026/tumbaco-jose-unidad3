# Sistema de Gestión de Libros Electrónicos

## Objetivo

Desarrollar un sistema de gestión de libros electrónicos utilizando Go,
SQLite y el patrón de acceso a datos DAO.

El proyecto permite almacenar y administrar información de libros
electrónicos mediante una entidad de modelo, una conexión a SQLite y
un objeto de acceso a datos (DAO).

## Tecnologías utilizadas

- Go
- SQLite
- database/sql
- modernc.org/sqlite
- Git y GitHub
- Visual Studio Code

## Estructura del proyecto

```text
deber4-docente_dao/
├── dao/
│   └── libro_electronico_dao.go
├── dataaccess/
│   └── dataaccess.go
├── model/
│   └── libro_electronico.go
├── main.go
├── go.mod
├── go.sum
├── .gitignore
├── README.md
└── BITACORA.md

Modelo

La entidad principal del proyecto es LibroElectronico.

Sus atributos son:

ID: identificador del libro.
Titulo: título del libro electrónico.
Autor: autor del libro.
Genero: género literario o temático.
ISBN: identificador internacional del libro.
Formato: formato digital del libro.
Precio: precio del libro electrónico.

La entidad también contiene los métodos:

MostrarInformacion()
ActualizarPrecio()
Acceso a datos

El paquete dataaccess se encarga de inicializar la conexión con
SQLite.

El paquete dao contiene LibroElectronicoDAO, encargado de realizar
las operaciones de persistencia sobre la tabla libros_electronicos.

Las operaciones implementadas son:

Crear la tabla.
Insertar un libro.
Obtener un libro por ID.
Obtener todos los libros.
Actualizar un libro.
Eliminar un libro.
Base de datos

La aplicación utiliza una base de datos SQLite llamada:

libros_electronicos.db

La tabla libros_electronicos contiene:

Campo	Tipo
id	INTEGER
titulo	TEXT
autor	TEXT
genero	TEXT
isbn	TEXT
formato	TEXT
precio	REAL

El campo id utiliza autoincremento.

El campo isbn se establece como único para evitar registrar dos
libros con el mismo ISBN.

El archivo libros_electronicos.db se excluye del repositorio mediante
.gitignore porque es una base de datos local generada durante la
ejecución.

Ejecución

Desde la carpeta del proyecto:

go fmt ./...
go run .

También se verificó la compilación mediante:

go test ./...
go build ./...

El comando go test ./... no reporta errores; actualmente no existen
archivos de pruebas automatizadas.

Evidencias de ejecución

Durante las pruebas se obtuvo:

Base de datos inicializada correctamente
Tabla de libros electrónicos creada correctamente
Libro creado correctamente con ID: 1
Libro encontrado: &{ID:1 Titulo:Clean Code Autor:Robert C. Martin Genero:Programación ISBN:9780132350884 Formato:PDF Precio:29.99}
Lista de libros:
ID: 1 | Título: Clean Code | Autor: Robert C. Martin | Género: Programación | ISBN: 9780132350884 | Formato: PDF | Precio: 29.99
Libro actualizado correctamente
Libro actualizado: &{ID:1 Titulo:Clean Code Autor:Robert C. Martin Genero:Programación ISBN:9780132350884 Formato:PDF Precio:39.99}
Libro eliminado correctamente

Estas evidencias demuestran:

Inicialización correcta de SQLite.
Creación de la tabla.
Inserción de un libro.
Consulta de un libro por ID.
Consulta de todos los libros.
Actualización del precio.
Eliminación del libro.
Decisiones de diseño

Se separó el proyecto en tres paquetes principales:

model: contiene la entidad LibroElectronico.
dataaccess: contiene la inicialización de la conexión SQLite.
dao: contiene las operaciones de persistencia.

Esta separación permite mantener independiente la representación de los
datos, la conexión con la base de datos y las operaciones de acceso a
datos.

Para los libros insertados se utiliza LastInsertId() para obtener el
identificador generado automáticamente por SQLite. Esto permite trabajar
con el ID real asignado por la base de datos.

El DAO recibe una conexión *sql.DB mediante:

NewLibroElectronicoDAO(db)

De esta manera, main.go se encarga de inicializar la base de datos y
el DAO se encarga de las operaciones sobre los libros.

Autor

José Tumbaco