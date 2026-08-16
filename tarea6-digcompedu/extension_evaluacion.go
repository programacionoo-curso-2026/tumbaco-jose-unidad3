package main

import (
	"fmt"
	"digcompedu-api/internal/domain"
)

// CalcularIndiceGlobal agrega un elemento de evaluación:
// un índice global de competencia digital docente.
// Se calcula como el promedio de los niveles de las 3 áreas evaluadas.
func CalcularIndiceGlobal(resultado domain.MatrizResultado) (float64, error) {
	if len(resultado.Resultados) == 0 {
		return 0, fmt.Errorf("no existen áreas evaluadas")
	}

	var suma int
	for _, area := range resultado.Resultados {
		if area.Nivel < domain.NivelInicial || area.Nivel > domain.NivelExperto {
			return 0, fmt.Errorf("nivel fuera de rango: %d", area.Nivel)
		}
		suma += int(area.Nivel)
	}

	return float64(suma) / float64(len(resultado.Resultados)), nil
}

func main() {
	resultado := domain.MatrizResultado{
		Resultados: []domain.AreaResultado{
			{Area: domain.AreaRecursosDigitales, Nivel: domain.NivelAvanzado},
			{Area: domain.AreaEvaluacion, Nivel: domain.NivelIntermedio},
			{Area: domain.AreaEmpoderamientoEstudiante, Nivel: domain.NivelAvanzado},
		},
	}

	indice, err := CalcularIndiceGlobal(resultado)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Índice global de competencia digital: %.2f/4.00\n", indice)
}
