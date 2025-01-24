package main

import "fmt"

func main() {
	// 0 1 1 2 3 5 8 13 21...
	indice := 25
	for i, j := uint(0), 0; i <= uint(indice); i++ {
		fmt.Println(j, fibonacci(i))
		j++
	}
}

func fibonacci(indice uint) uint {
	if indice <= 1 {
		return indice
	}
	return fibonacci(indice-2) + fibonacci(indice-1)
}
