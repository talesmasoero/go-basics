package main

import "fmt"

func main() {
	str := "Tales"
	switch {
	case str == "Tales":
		fmt.Println("Primeiro case")
		fallthrough
	case len(str) == 5:
		fmt.Println("Segundo case")
		fallthrough
	default:
		fmt.Println("Default")
	}
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(-10))

	n := teste(10)
	fmt.Println(n)
}

func diaDaSemana(d int) string {
	switch d {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	}
	return "Número inválido"
}

func teste(n int) int {
	switch n {
	case 10:
		n = 1
		fallthrough
	case 100:
		n = 5
	}
	return n
}
