package dao

import (
	"database/sql"
	"deber4-docente_dao/model"
)

type LibroElectronicoDAO struct {
	DB *sql.DB
}

func NewLibroElectronicoDAO(db *sql.DB) *LibroElectronicoDAO {
	return &LibroElectronicoDAO{
		DB: db,
	}
}

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

	_, err := dao.DB.Exec(query)
	return err
}

func (dao *LibroElectronicoDAO) Create(libro model.LibroElectronico) (int64, error) {
	query := `
	INSERT INTO libros_electronicos
	(titulo, autor, genero, isbn, formato, precio)
	VALUES (?, ?, ?, ?, ?, ?);
	`

	result, err := dao.DB.Exec(
		query,
		libro.Titulo,
		libro.Autor,
		libro.Genero,
		libro.ISBN,
		libro.Formato,
		libro.Precio,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (dao *LibroElectronicoDAO) GetByID(id int64) (*model.LibroElectronico, error) {
	query := `
	SELECT id, titulo, autor, genero, isbn, formato, precio
	FROM libros_electronicos
	WHERE id = ?;
	`

	libro := &model.LibroElectronico{}

	err := dao.DB.QueryRow(query, id).Scan(
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

func (dao *LibroElectronicoDAO) GetAll() ([]model.LibroElectronico, error) {
	query := `
	SELECT id, titulo, autor, genero, isbn, formato, precio
	FROM libros_electronicos
	ORDER BY id;
	`

	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []model.LibroElectronico

	for rows.Next() {
		var libro model.LibroElectronico

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

func (dao *LibroElectronicoDAO) Update(libro model.LibroElectronico) error {
	query := `
	UPDATE libros_electronicos
	SET titulo = ?, autor = ?, genero = ?, isbn = ?, formato = ?, precio = ?
	WHERE id = ?;
	`

	_, err := dao.DB.Exec(
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

func (dao *LibroElectronicoDAO) Delete(id int64) error {
	query := `
	DELETE FROM libros_electronicos
	WHERE id = ?;
	`

	_, err := dao.DB.Exec(query, id)
	return err
}
