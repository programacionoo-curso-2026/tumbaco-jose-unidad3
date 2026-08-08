# Taller 2 - Data Access con Go y SQLite

## Objetivo

Desarrollar una aplicación en Go que permita gestionar información de videojuegos mediante una entidad de modelo y una capa de acceso a datos utilizando SQLite.

El proyecto permite realizar operaciones de persistencia sobre videojuegos mediante el paquete `database/sql`.

## Tecnologías utilizadas

- Go
- SQLite
- `database/sql`
- `modernc.org/sqlite`
- Git y GitHub
- Visual Studio Code

## Estructura del proyecto

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
├── README.md
└── BITACORA.md
```

## Modelo

La entidad principal del proyecto es `Videojuego`.

Sus atributos son:

- `ID`: identificador del videojuego.
- `Nombre`: nombre del videojuego.
- `Genero`: género del videojuego.
- `Plataforma`: plataforma en la que se ejecuta.
- `Precio`: precio del videojuego.

La entidad también contiene los siguientes métodos:

- `MostrarInformacion()`
- `ActualizarPrecio()`

## Acceso a datos

El paquete `dataaccess` se encarga de la conexión y las operaciones con SQLite.

Las operaciones implementadas son:

| Operación | Método | Descripción |
|---|---|---|
| Crear tabla | `CrearTabla()` | Crea la tabla `videojuegos` si no existe. |
| Crear | `CrearVideojuego()` | Inserta un nuevo videojuego y devuelve el ID generado. |
| Consultar uno | `ObtenerVideojuegoPorID()` | Busca un videojuego utilizando su ID. |
| Consultar todos | `ObtenerVideojuegos()` | Recupera todos los videojuegos almacenados. |
| Actualizar | `ActualizarPrecio()` | Modifica el precio de un videojuego. |
| Eliminar | `EliminarVideojuego()` | Elimina un videojuego por su ID. |
| Cerrar conexión | `Close()` | Cierra la conexión con SQLite. |

Estas operaciones permiten realizar el ciclo principal de persistencia de datos: crear, consultar, actualizar y eliminar registros.

## Base de datos

La aplicación utiliza una base de datos SQLite llamada `videojuegos.db`.

La tabla `videojuegos` contiene los siguientes campos:

| Campo | Tipo | Descripción |
|---|---|---|
| `id` | INTEGER | Identificador único generado automáticamente. |
| `nombre` | TEXT | Nombre del videojuego. |
| `genero` | TEXT | Género del videojuego. |
| `plataforma` | TEXT | Plataforma del videojuego. |
| `precio` | REAL | Precio del videojuego. |

El archivo `videojuegos.db` se excluye del repositorio mediante `.gitignore` porque es una base de datos local generada durante la ejecución.

## Ejecución

Para ejecutar el proyecto, primero se puede formatear el código:

```powershell
go fmt ./...
```

Luego se ejecuta la aplicación:

```powershell
go run .
```

El programa crea la tabla si todavía no existe y ejecuta pruebas de las operaciones de acceso a datos.

## Pruebas realizadas

Durante el desarrollo se comprobó:

1. Conexión correcta con SQLite.
2. Creación de la tabla.
3. Inserción de videojuegos.
4. Consulta de un videojuego por ID.
5. Consulta de todos los videojuegos.
6. Actualización del precio.
7. Consulta del videojuego actualizado.
8. Eliminación de un videojuego.

Una ejecución final permitió comprobar correctamente las operaciones de creación, consulta, actualización y eliminación.

## Manejo del ID generado por SQLite

Durante las pruebas se detectó un problema al intentar consultar siempre el videojuego utilizando el ID `1`.

Después de eliminar registros anteriores, el ID `1` ya no correspondía a un videojuego existente, por lo que la consulta produjo el error:

```text
sql: no rows in result set
```

Para solucionar este problema, `CrearVideojuego()` utiliza `LastInsertId()` para obtener el identificador generado por SQLite después de insertar el registro.

De esta forma, el programa utiliza el ID real generado por la base de datos para realizar posteriormente las operaciones de consulta, actualización y eliminación.

## Decisiones de diseño

Se separó el proyecto en dos paquetes principales:

- `model`: contiene la entidad `Videojuego` y sus métodos.
- `dataaccess`: contiene la lógica relacionada con SQLite y las operaciones de persistencia.

Esta separación permite mantener diferenciados el modelo de datos y la lógica de acceso a la base de datos.

También se utiliza `database/sql` para trabajar con SQLite y consultas parametrizadas mediante `?`, evitando colocar directamente los valores dentro de las consultas SQL.

## Resultados

Al finalizar las pruebas se verificó que la aplicación puede:

- Conectarse correctamente a SQLite.
- Crear la tabla `videojuegos`.
- Insertar videojuegos.
- Obtener videojuegos por ID.
- Obtener todos los videojuegos.
- Actualizar precios.
- Eliminar videojuegos.
- Cerrar correctamente la conexión con la base de datos.

## Autor

José Tumbaco
