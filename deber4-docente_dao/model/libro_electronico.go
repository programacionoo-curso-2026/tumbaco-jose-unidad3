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
	return l.Titulo + " - " + l.Autor + " - " + l.Formato
}

func (l *LibroElectronico) ActualizarPrecio(nuevoPrecio float64) {
	l.Precio = nuevoPrecio
}
