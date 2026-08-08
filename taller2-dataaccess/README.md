\# Taller 2 - Data Access con Go y SQLite



\## Objetivo



Desarrollar una aplicación en Go que permita gestionar información de videojuegos mediante una entidad de modelo y una capa de acceso a datos utilizando SQLite.



El proyecto permite realizar operaciones de persistencia sobre videojuegos mediante el paquete `database/sql`.



\## Tecnologías utilizadas



\- Go

\- SQLite

\- `database/sql`

\- `modernc.org/sqlite`

\- Git y GitHub

\- Visual Studio Code



\## Estructura del proyecto



```text

taller2-dataaccess/

├── dataaccess/

│   └── dataaccess.go

├── model/

│   └── videojuego.go

├── main.go

├── go.mod

├── go.sum

├── .gitignore

└── README.md



Modelo



La entidad principal del proyecto es Videojuego.



Sus atributos son:



ID: identificador del videojuego.

Nombre: nombre del videojuego.

Genero: género del videojuego.

Plataforma: plataforma en la que se ejecuta.

Precio: precio del videojuego.



La entidad también contiene los métodos:



MostrarInformacion()

ActualizarPrecio()

Acceso a datos



El paquete dataaccess se encarga de la conexión y las operaciones con SQLite.



Las operaciones implementadas son:



Crear la tabla videojuegos.

Insertar un videojuego.

Obtener un videojuego por ID.

Obtener todos los videojuegos.

Actualizar el precio de un videojuego.

Eliminar un videojuego.

Cerrar la conexión con la base de datos.

Base de datos



La aplicación utiliza una base de datos SQLite llamada videojuegos.db.



La tabla videojuegos contiene:



Campo	Tipo

id	INTEGER

nombre	TEXT

genero	TEXT

plataforma	TEXT

precio	REAL



El archivo videojuegos.db se excluye del repositorio mediante .gitignore porque es una base de datos local generada durante la ejecución.



Ejecución



Para ejecutar el proyecto:



go fmt ./...

go run .



El programa crea la tabla si todavía no existe y ejecuta pruebas de las operaciones de acceso a datos.



Pruebas realizadas



Durante el desarrollo se comprobó:



Conexión correcta con SQLite.

Creación de la tabla.

Inserción de videojuegos.

Consulta de un videojuego por ID.

Consulta de todos los videojuegos.

Actualización del precio.

Eliminación de un videojuego.

Decisiones de diseño



Se separó el proyecto en dos paquetes principales:



model: contiene la entidad Videojuego y sus métodos.

dataaccess: contiene la lógica relacionada con SQLite y las operaciones de persistencia.



Para los videojuegos insertados se utiliza LastInsertId() para obtener el ID generado por SQLite. Esto evita depender de un ID fijo al realizar consultas, actualizaciones o eliminaciones.



Autor



José Tumbaco

