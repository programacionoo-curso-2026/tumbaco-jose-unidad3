package model

type LibroElectronico struct {
	ID      int
	Titulo  string
	Autor   string
	Genero  string
	ISBN    string
	Formato string
	Precio  float64
}

func (l LibroElectronico) MostrarInformacion() string {
	return "Título: " + l.Titulo +
		" | Autor: " + l.Autor +
		" | Género: " + l.Genero +
		" | ISBN: " + l.ISBN +
		" | Formato: " + l.Formato
}

func (l *LibroElectronico) ActualizarPrecio(nuevoPrecio float64) {
	l.Precio = nuevoPrecio
}
