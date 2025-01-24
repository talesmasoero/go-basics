package main

import (
	"fmt"
	"time"
)

// A função principal escreve há quantos segundos
// o programa está sendo executado
func main() {
	fmt.Println("Iniciando execução...")

	// Como essa função é iniciada com uma goroutine, ela executa
	// concorrentemente com a função principal e escreve GOROUTINE
	// a cada 3/4 de segundo
	go func() {
		for {
			fmt.Println("GOROUTINE")
			time.Sleep(time.Millisecond * 750)
		}
	}()

	// "Função Principal"
	sec := "segundo"
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i != 1 {
			sec = "segundos"
		}
		fmt.Printf("O progama está sendo executado há %d %s\n", i, sec)
	}
	fmt.Println("Encerrando execução...")
}
