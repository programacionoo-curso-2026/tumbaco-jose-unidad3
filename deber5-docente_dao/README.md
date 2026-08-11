# Sistema de Gestión de Libros Electrónicos — Deber 5

## 1. Descripción

Este proyecto implementa un sistema de acceso a datos para gestionar libros electrónicos utilizando **Go**, **SQLite** y el patrón **DAO (Data Access Object)**.

El proyecto está organizado en diferentes paquetes para separar las responsabilidades del sistema:

* `model`: contiene la entidad `LibroElectronico`.
* `dataaccess`: administra la conexión con la base de datos SQLite.
* `dao`: contiene las operaciones de acceso a datos.
* `main.go`: ejecuta y demuestra las operaciones CRUD.

## 2. Tecnologías utilizadas

* Go
* SQLite
* `database/sql`
* `modernc.org/sqlite`
* Patrón DAO

## 3. Estructura del proyecto

```text
deber5-docente_dao/
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
└── libros_electronicos.db
```

La base de datos `libros_electronicos.db` es un archivo local y está excluida del control de versiones mediante `.gitignore`.

## 4. Modelo de datos

La entidad `LibroElectronico` contiene los siguientes atributos:

| Campo   | Tipo    |
| ------- | ------- |
| ID      | int     |
| Titulo  | string  |
| Autor   | string  |
| Genero  | string  |
| ISBN    | string  |
| Formato | string  |
| Precio  | float64 |

También se implementaron métodos para mostrar información del libro y actualizar su precio.

## 5. Operaciones DAO

El `LibroElectronicoDAO` implementa las siguientes operaciones:

* `CreateTable()` — crea la tabla `libros_electronicos`.
* `Create()` — inserta un nuevo libro.
* `GetByID()` — busca un libro mediante su ID.
* `GetAll()` — obtiene todos los libros.
* `Update()` — actualiza la información de un libro.
* `Delete()` — elimina un libro mediante su ID.

## 6. Base de datos

La aplicación utiliza SQLite mediante el paquete:

```text
modernc.org/sqlite
```

La conexión se inicializa desde el paquete `dataaccess`.

El archivo local utilizado por la aplicación es:

```text
libros_electronicos.db
```

Este archivo no se incluye en Git gracias a las reglas definidas en `.gitignore`.

## 7. Ejecución del programa

Para ejecutar el proyecto se utiliza:

```bash
go run .
```

Durante la prueba se obtuvo la siguiente evidencia:

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

## 8. Evidencias de funcionamiento

La ejecución permitió comprobar correctamente las principales operaciones del sistema:

### Crear

Se creó correctamente el libro:

```text
Título: Clean Code
Autor: Robert C. Martin
Género: Programación
ISBN: 9780132350884
Formato: PDF
Precio: 29.99
```

El sistema asignó automáticamente el ID:

```text
1
```

### Leer

Se realizó correctamente la búsqueda del libro mediante su ID:

```text
Libro encontrado: &{ID:1 Titulo:Clean Code Autor:Robert C. Martin Genero:Programación ISBN:9780132350884 Formato:PDF Precio:29.99}
```

### Leer todos

Se obtuvo correctamente la lista de libros registrados:

```text
ID: 1 | Título: Clean Code | Autor: Robert C. Martin | Género: Programación | ISBN: 9780132350884 | Formato: PDF | Precio: 29.99
```

### Actualizar

Se modificó correctamente el precio:

```text
29.99 → 39.99
```

La actualización fue comprobada mediante una nueva consulta.

### Eliminar

Finalmente, el libro fue eliminado correctamente:

```text
Libro eliminado correctamente
```

## 9. Verificación del código

También se ejecutaron los comandos:

```bash
go fmt ./...
```

y:

```bash
go test ./...
```

Los paquetes del proyecto compilaron correctamente y no se reportaron errores.

## 10. Control de versiones

El proyecto fue agregado al repositorio Git mediante un commit incremental:

```text
2bd9f43 deber5: implementa sistema de libros electronicos
```

El archivo de base de datos local no fue incluido en el commit debido a las reglas establecidas en `.gitignore`.

## 11. Conclusión

El programa permite administrar libros electrónicos mediante una arquitectura organizada en modelo, acceso a datos y DAO.

La ejecución comprobó correctamente la creación de la tabla y las operaciones de **crear, consultar, listar, actualizar y eliminar** registros utilizando SQLite.
