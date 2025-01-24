package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Goroutine para enviar mensagem ao canal1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// Goroutine para enviar mensagem ao canal2
	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			canal2 <- "Canal 2"
		}
	}()

	// Usando o select para receber de ambos os canais
	for {
		select {
		case msg := <-canal1:
			fmt.Println(msg)
		case msg := <-canal2:
			fmt.Println(msg)
		}
	}
}
