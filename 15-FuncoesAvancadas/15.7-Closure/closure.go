package main

import (
	"fmt"
	"time"
)

func main() {
	// Criação das instâncias das funções closure
	increaseSlice1 := createIntSlice()
	increaseSlice2 := createIntSlice()

	// Laço para simular o crescimento diferenciado
	go func() {
		var i int
		for i < 10 {
			fmt.Println("Slice 1:", increaseSlice1())
			time.Sleep(1000 * time.Millisecond) // Slice 1 cresce mais rápido
			i++
		}
	}()

	go func() {
		var i int
		for i < 10 {
			fmt.Println("Slice 2:", increaseSlice2())
			time.Sleep(2000 * time.Millisecond) // Slice 2 cresce mais devagar
			i++
		}
	}()

	// Mantém o programa rodando
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Os slices são independentes entre si e crescem em velocidades diferentes.")
}

// Função Closure
func createIntSlice() func() []int {
	intSlice := []int{}
	i := 0

	return func() []int {
		intSlice = append(intSlice, i)
		i++
		return intSlice
	}
}
