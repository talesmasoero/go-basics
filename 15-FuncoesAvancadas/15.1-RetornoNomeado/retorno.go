package main

import "fmt"

func main() {
	soma, sub, mult, div := operacoes(10, 5)
	fmt.Println(soma, sub, mult, div)
}

func operacoes(n1, n2 int) (soma, sub, mult, div int) {
	soma = n1 + n2
	sub = n1 - n2
	mult = n1 * n2
	div = n1 / n2
	return
}
