# Bitácora de avances — Deber 5

## 1. Creación del proyecto

Se creó la carpeta:

```text
deber5-docente_dao
```

El proyecto se configuró como un módulo de Go mediante:

```text
go mod init deber5-docente_dao
```

El módulo quedó configurado con Go 1.26.3.

## 2. Organización del proyecto

Se creó una estructura basada en la separación de responsabilidades:

```text
deber5-docente_dao/
├── dao/
├── dataaccess/
├── model/
├── main.go
├── go.mod
├── go.sum
└── .gitignore
```

La implementación utiliza el patrón DAO para separar el acceso a datos de la lógica principal.

## 3. Configuración de SQLite

Se agregó la dependencia:

```text
modernc.org/sqlite
```

mediante:

```text
go get modernc.org/sqlite
```

Posteriormente se ejecutó:

```text
go mod tidy
```

sin errores.

La conexión a SQLite se implementó en:

```text
dataaccess/dataaccess.go
```

La aplicación utiliza la base de datos local:

```text
libros_electronicos.db
```

## 4. Modelo de libro electrónico

Se implementó:

```text
model/libro_electronico.go
```

con la entidad `LibroElectronico`.

El modelo contiene:

* ID
* Titulo
* Autor
* Genero
* ISBN
* Formato
* Precio

También se implementaron métodos para mostrar información y actualizar el precio del libro.

## 5. Implementación del DAO

Se creó:

```text
dao/libro_electronico_dao.go
```

El DAO implementa las siguientes operaciones:

* `CreateTable()`
* `Create()`
* `GetByID()`
* `GetAll()`
* `Update()`
* `Delete()`

De esta manera se implementaron las operaciones principales de acceso a datos utilizando SQLite.

## 6. Verificación de compilación

Se ejecutó:

```text
go fmt ./...
```

para formatear el código.

Posteriormente se ejecutó:

```text
go test ./...
```

La salida confirmó que los paquetes compilaban correctamente:

```text
?       deber5-docente_dao/dao        [no test files]
?       deber5-docente_dao/dataaccess  [no test files]
?       deber5-docente_dao/model       [no test files]
```

No se presentaron errores de compilación.

## 7. Implementación de main.go

Se desarrolló `main.go` para demostrar el funcionamiento del sistema.

El programa realiza las siguientes operaciones:

1. Inicializa la base de datos.
2. Crea la tabla de libros electrónicos.
3. Inserta un libro.
4. Busca el libro por ID.
5. Obtiene todos los libros.
6. Actualiza el precio.
7. Comprueba la actualización.
8. Elimina el libro.

Para la prueba se utilizó:

```text
Título: Clean Code
Autor: Robert C. Martin
Género: Programación
ISBN: 9780132350884
Formato: PDF
Precio inicial: 29.99
```

## 8. Evidencia de ejecución

Se ejecutó:

```text
go run .
```

La ejecución fue exitosa y produjo:

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

## 9. Verificación de operaciones

La ejecución permitió comprobar:

### Create

Se insertó correctamente el libro `Clean Code` y SQLite asignó el ID `1`.

### Read

Se recuperó correctamente el libro mediante su ID.

### Read All

Se obtuvo correctamente la lista de libros registrados.

### Update

El precio se modificó de:

```text
29.99
```

a:

```text
39.99
```

La modificación fue comprobada mediante una nueva consulta.

### Delete

El registro fue eliminado correctamente.

## 10. Corrección durante el desarrollo

Durante el desarrollo inicial se interpretó temporalmente que el proyecto debía implementar un DAO para docentes debido al nombre `deber5-docente_dao` y a la referencia de la consigna.

Se llegaron a crear temporalmente archivos relacionados con `Docente`, pero antes de realizar commits se identificó que el objetivo de trabajo establecido era continuar con el mismo **Sistema de Gestión de Libros Electrónicos** desarrollado anteriormente.

Los archivos relacionados con docentes fueron eliminados antes de registrar los cambios en Git.

Posteriormente se reconstruyó correctamente el proyecto utilizando:

```text
LibroElectronico
LibroElectronicoDAO
libros_electronicos.db
```

y se volvió a comprobar la ejecución completa del CRUD sin errores.

## 11. Protección de la base de datos

Se creó `.gitignore` con reglas para evitar subir archivos locales de SQLite:

```text
*.db
*.db-shm
*.db-wal
```

De esta manera, `libros_electronicos.db` permanece fuera del repositorio.

## 12. Control de versiones

Se realizó un primer commit con la implementación principal:

```text
2bd9f43 deber5: implementa sistema de libros electronicos
```

Posteriormente se agregó la documentación del proyecto mediante:

```text
d736fcc deber5: documenta ejecucion y evidencias
```

Los commits se realizaron desde el repositorio principal `tumbaco-jose-unidad3`.

## 13. Estado actual

El proyecto cuenta actualmente con:

* Modelo `LibroElectronico`.
* Conexión SQLite.
* `LibroElectronicoDAO`.
* Operaciones CRUD.
* `main.go` funcional.
* `.gitignore`.
* `README.md`.
* `BITACORA.md`.
* Evidencias de ejecución documentadas.
* Historial de commits incremental.

La ejecución final del programa fue exitosa y permitió comprobar el funcionamiento de las operaciones de creación, consulta, listado, actualización y eliminación de libros electrónicos.
