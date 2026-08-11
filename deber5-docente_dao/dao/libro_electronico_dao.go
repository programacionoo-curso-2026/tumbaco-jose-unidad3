package dao

import (
	"database/sql"

	"deber5-docente_dao/model"
)

type LibroElectronicoDAO struct {
	db *sql.DB
}

func NewLibroElectronicoDAO(db *sql.DB) *LibroElectronicoDAO {
	return &LibroElectronicoDAO{db: db}
}

// CreateTable crea la tabla de libros electrónicos.
func (dao *LibroElectronicoDAO) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS libros_electronicos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT NOT NULL,
		autor TEXT NOT NULL,
		genero TEXT NOT NULL,
		isbn TEXT NOT NULL UNIQUE,
		formato TEXT NOT NULL,
		precio REAL NOT NULL
	);
	`

	_, err := dao.db.Exec(query)
	return err
}

// Create inserta un nuevo libro electrónico.
func (dao *LibroElectronicoDAO) Create(libro *model.LibroElectronico) error {
	query := `
	INSERT INTO libros_electronicos
		(titulo, autor, genero, isbn, formato, precio)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := dao.db.Exec(
		query,
		libro.Titulo,
		libro.Autor,
		libro.Genero,
		libro.ISBN,
		libro.Formato,
		libro.Precio,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	libro.ID = int(id)

	return nil
}

// GetByID busca un libro por su ID.
func (dao *LibroElectronicoDAO) GetByID(id int) (*model.LibroElectronico, error) {
	query := `
	SELECT id, titulo, autor, genero, isbn, formato, precio
	FROM libros_electronicos
	WHERE id = ?
	`

	libro := &model.LibroElectronico{}

	err := dao.db.QueryRow(query, id).Scan(
		&libro.ID,
		&libro.Titulo,
		&libro.Autor,
		&libro.Genero,
		&libro.ISBN,
		&libro.Formato,
		&libro.Precio,
	)
	if err != nil {
		return nil, err
	}

	return libro, nil
}

// GetAll obtiene todos los libros electrónicos.
func (dao *LibroElectronicoDAO) GetAll() ([]*model.LibroElectronico, error) {
	query := `
	SELECT id, titulo, autor, genero, isbn, formato, precio
	FROM libros_electronicos
	ORDER BY id
	`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []*model.LibroElectronico

	for rows.Next() {
		libro := &model.LibroElectronico{}

		err := rows.Scan(
			&libro.ID,
			&libro.Titulo,
			&libro.Autor,
			&libro.Genero,
			&libro.ISBN,
			&libro.Formato,
			&libro.Precio,
		)
		if err != nil {
			return nil, err
		}

		libros = append(libros, libro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return libros, nil
}

// Update actualiza un libro electrónico.
func (dao *LibroElectronicoDAO) Update(libro *model.LibroElectronico) error {
	query := `
	UPDATE libros_electronicos
	SET titulo = ?,
		autor = ?,
		genero = ?,
		isbn = ?,
		formato = ?,
		precio = ?
	WHERE id = ?
	`

	_, err := dao.db.Exec(
		query,
		libro.Titulo,
		libro.Autor,
		libro.Genero,
		libro.ISBN,
		libro.Formato,
		libro.Precio,
		libro.ID,
	)

	return err
}

// Delete elimina un libro electrónico por su ID.
func (dao *LibroElectronicoDAO) Delete(id int) error {
	query := `
	DELETE FROM libros_electronicos
	WHERE id = ?
	`

	_, err := dao.db.Exec(query, id)
	return err
}
