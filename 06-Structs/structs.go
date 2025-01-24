package main

import (
	s "basics/06-Structs/structsaux"
	"fmt"
)

func main() {
	p1 := s.Pessoa{
		Nome:  "Tales",
		Idade: 19,
		Endereco: s.Endereco{
			Cidade: "BSB",
			Numero: 4,
		},
	}

	var p2 s.Pessoa
	p2.Nome = "Breno"
	p2.Idade = 14
	p2.Endereco.Cidade = "BSB"
	p2.Endereco.Numero = 4

	p3 := s.Pessoa{Nome: "Eduardo", Idade: 20, Endereco: s.Endereco{Cidade: "BSB"}}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
