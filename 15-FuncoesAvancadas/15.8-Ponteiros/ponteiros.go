package main

import "fmt"

func main() {
	// Passando cópias
	array1 := [3]int{1, 2, 3}
	arrayCopia := recebeCopia(array1)

	fmt.Println("Array original (não tem seu valor altereado):", array1)
	fmt.Println("'Cópia' retornada:", arrayCopia)

	// Passando ponteiros
	array2 := [3]int{4, 5, 6}
	fmt.Println("Array original (antes da função):", array2)

	// A função não tem retorno, apenas altera o valor de array2
	recebePonteiro(&array2)
	fmt.Println("Array original (depois da função):", array2)
}

// As funções resetam os elementos do array para 0

// arr é passado como cópia de array
func recebeCopia(arr [3]int) [3]int {
	for i := range arr {
		arr[i] = 0
	}
	return arr
}

// arr é passado como ponteiro (endereço de memória)
func recebePonteiro(arr *[3]int) {
	for i := range *arr {
		arr[i] = 0
	}
}
