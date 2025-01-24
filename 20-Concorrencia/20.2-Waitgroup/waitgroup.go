package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("GOROUTINE 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("GOROUTINE 2")
		waitGroup.Done()
	}()

	// wg.Wait() impede que o programa continue até
	// que as goroutines sejam finalizadas.
	waitGroup.Wait()

	// Só vai ser executado depois que as goroutines
	// sejam finalizadas
	fmt.Println("Tive que esperar o WaitGroup...")
}

func escrever(str string) {
	for i := range 5 {
		fmt.Println(i, str)
		time.Sleep(time.Millisecond * 500)
	}
}
