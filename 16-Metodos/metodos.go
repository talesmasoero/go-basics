package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) adicionarIdade(years int) *Pessoa {
	p.Idade += years
	return p
}

func (p *Pessoa) trocarNome(newNome string) *Pessoa {
	p.Nome = newNome
	return p
}

func main() {
	p := &Pessoa{Nome: "Breno", Idade: 15}
	// Encadeando métodos (só é possível porque os métodos retornam o tipo correto)
	p.adicionarIdade(5).trocarNome("Tales")
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
}
