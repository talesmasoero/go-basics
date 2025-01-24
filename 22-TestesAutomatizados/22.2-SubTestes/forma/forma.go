package forma

import "math"

type Forma interface {
	Area() float64
}

type Quadrado struct {
	Lado float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Altura int
	Base   int
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}

func (t Triangulo) Area() int {
	return (t.Base * t.Altura) / 2
}
