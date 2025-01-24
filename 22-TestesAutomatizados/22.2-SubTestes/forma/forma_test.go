package forma

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Quadrado", func(t *testing.T) {
		quadrado := Quadrado{Lado: 10}
		areaEsperada := float64(100) // Lado ao quadrado
		areaRecebida := quadrado.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("\nEsperado: %.2f\nRecebido: %.2f\n", areaEsperada, areaRecebida)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{Raio: 5}
		areaEsperada := math.Pi * 25 // Pi vezes o raio ao quadrado
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("\nEsperado: %.2f\nRecebido: %.2f\n", areaEsperada, areaRecebida)
		}
	})

	t.Run("Triangulo", func(t *testing.T) {
		triangulo := Triangulo{Altura: 4, Base: 3}
		areaEsperada := 6 // Base vezes altura sobre dois
		areaRecebida := triangulo.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("\nEsperado: %d\nRecebido: %d\n", areaEsperada, areaRecebida)
		}
	})
}
