package main

import "math"

type Forma interface {
	// Assinatura da função
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

// Funções que tem a assinatura implementada pela interface
func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}

// Função que não tem a assinatura implementada pela interface
func (t Triangulo) Area() int {
	return (t.Base * t.Altura) / 2
}
