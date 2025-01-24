package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade uint8
}

type Estudante struct {
	Curso string
	Pessoa
}

func main() {
	var p Pessoa

	p.Nome = "Tales"
	p.Idade = 19

	e1 := Estudante{Curso: "CC", Pessoa: p}

	var e2 Estudante
	e2.Nome = "Eduardo"
	e2.Idade = 20
	e2.Curso = "ADM"

	e3 := Estudante{
		Curso: "",
		Pessoa: Pessoa{
			Nome:  "Breno",
			Idade: 14,
		},
	}

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
