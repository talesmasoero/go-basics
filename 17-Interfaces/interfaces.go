package main

import "fmt"

func main() {
	quadrado := Quadrado{
		Lado: 8,
	}

	circulo := Circulo{
		Raio: 5,
	}

	triangulo := Triangulo{
		Altura: 5,
		Base:   10,
	}

	// Erro: cannot use triangulo (variable of type Triangulo) as Forma value in array or slice literal:
	// Triangulo does not implement Forma (wrong type for method Area)
	// have Area() int
	// want Area() float64 (compilerInvalidIfaceAssign)
	//
	// formas := []Forma{quadrado, circulo, triangulo}

	formas := []Forma{quadrado, circulo /*, triangulo*/}
	for _, f := range formas {
		fmt.Printf("Essa forma geométrica tem a área de: %.2fm\n", f.Area())
	}

	fmt.Printf("Triângulo (não implementa a interface): %.dm\n", triangulo.Area())
}
