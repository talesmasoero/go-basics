package main

import "fmt"

func main() {
	mapChars := make(map[rune]int)
	str := "arara"
	// Contando a quantidade de cada caractere na string
	// a, r = 3, 2
	for _, r := range str {
		mapChars[r]++
	}
	if qtd, ok := mapChars['a']; ok {
		fmt.Printf("Há %d a's em %s\n", qtd, str)
	}
	if qtd, ok := mapChars['t']; ok {
		fmt.Printf("Há %d t's em %s\n", qtd, str)
	} else {
		fmt.Printf("Não há t's em %s\n", str)
	}

	idade := 20
	if idade < 13 {
		fmt.Println("Criança")
	} else if idade >= 13 && idade < 18 {
		fmt.Println("Adolescente")
	} else {
		fmt.Println("Maior de idade")
	}
}
