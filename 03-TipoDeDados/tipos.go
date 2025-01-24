package main

import (
	"fmt"
)

func main() {
	str := "Hello, World"
	fmt.Printf("Char\tIndex\tValue\tType\n")
	for i, r := range str {
		fmt.Printf("%s\t%d\t%d\t%T\n", string(r), i, r, r)
		if r == 'o' {
			fmt.Println()
		}
	}
}
