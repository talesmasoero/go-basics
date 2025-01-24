package main

import "fmt"

func main() {
	var ptrInt *int
	var num int

	ptrInt = &num
	num = 10

	fmt.Printf("\tEndereço\tDado\nNum:\t%p\t%d\nPtr:\t%p\t%p\n\n", &num, num, &ptrInt, ptrInt)
	fmt.Printf("O dado contido em ptrInt é o endereço de num,\nque por sua vez guarda o dado: %d\n\n", *ptrInt)

	*ptrInt++

	fmt.Printf("Alterando o ponteiro, também se altera a variável: %d, %d", num, *ptrInt)
}
