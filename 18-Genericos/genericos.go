package main

import "fmt"

type Inteiros interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Reais interface {
	~float32 | ~float64
}

type Numeros interface {
	Inteiros | Reais
}

func main() {
	inteiros1 := []int16{1, 2, 3, 4, 5}
	inteiros2 := []int64{6, 7, 8, 9, 0}
	reais1 := []float32{1.1, 1.2, 1.3}
	reais2 := []float64{1.4, 1.5, 1.6}

	fmt.Println(somaNumeros(inteiros1))
	fmt.Println(somaNumeros(inteiros2))
	fmt.Println(somaNumeros(reais1))
	fmt.Println(somaNumeros(reais2))
}

// Em vez de criarmos uma função para somar cada tipo de slice,
// criamos apenas uma com tipos genéricos
func somaNumeros[T Numeros](numeros []T) T {
	var soma T
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
