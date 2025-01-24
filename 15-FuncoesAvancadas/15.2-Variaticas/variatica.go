package main

import "fmt"

func main() {
	soma("Hello, World!", 10, 2, 2, 100, 32, 25, 120, 80, -6)
}

func soma(str string, numeros ...int) {
	var total int
	for _, num := range numeros {
		total += num
	}
	fmt.Println(str, total)
}
