package main

import (
	"fmt"
	"time"
)

func main() {
	exemplo1()
	fmt.Println("======================================")
	exemplo2()
}

func escrever(str string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Goroutine %d: %s", i, str)
		// O código trava aqui até que seja retirado no loop
		fmt.Println("Teste de travamento")
	}
	close(ch)
}

func exemplo1() {
	fmt.Println("Antes da goroutine...")

	ch := make(chan string)
	// "Popula" o canal por meio de uma goroutine
	go escrever("Hello, world!", ch)

	fmt.Println("Depois da goroutine começar a execução...")

	// Ao printarmos os itens que estão no canal, também o retiramos dele
	for str := range ch {
		// O código da goroutine é travado a cada iteração desse loop
		fmt.Println(str)
		time.Sleep(time.Millisecond * 2000)
	}

	_, open2 := <-ch
	fmt.Println("Canal aberto:", open2)
}

func exemplo2() {
	ch := make(chan int)

	fmt.Println("Início do programa.")

	go func() {
		fmt.Println("Código antes do sleep...")
		time.Sleep(time.Millisecond * 5000)
		ch <- 50000
		fmt.Println("Número enviado...")
	}()

	fmt.Println("Código após a goroutine...")

	num := <-ch
	fmt.Println("Número recebeido:", num)
	fmt.Println("Final do programa.")
}
