package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	i := 10
	for i <= 50 {
		ch <- i
		i += 10
	}
	close(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
