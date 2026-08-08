package model

type Videojuego struct {
	ID         int
	Nombre     string
	Genero     string
	Plataforma string
	Precio     float64
}

func (v Videojuego) MostrarInformacion() string {
	return v.Nombre + " - " + v.Genero + " - " + v.Plataforma
}

func (v *Videojuego) ActualizarPrecio(nuevoPrecio float64) {
	v.Precio = nuevoPrecio
}
