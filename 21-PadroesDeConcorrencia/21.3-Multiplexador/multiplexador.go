package main

import (
	"fmt"
	"time"
)

func main() {
	ch := multiplexer(print("Channel 1", 1), print("Channel 2", 2))
	for msg := range ch {
		fmt.Println(msg)
	}
}

func multiplexer(ch1, ch2 <-chan string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				ch <- msg
			case msg := <-ch2:
				ch <- msg
			}
		}
	}()

	return ch
}

// Retorna um canal de leitura
func print(str string, sec time.Duration) <-chan string {
	ch := make(chan string)

	go func() {
		for range 10 {
			time.Sleep(time.Second * sec)
			ch <- str
		}
	}()

	return ch
}
