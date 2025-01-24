package main

import "fmt"

func main() {
	slc := make([]int, 10, 15)
	fmt.Println(slc, len(slc), cap(slc))

	for i := 0; i < 6; i++ {
		slc = append(slc, 0)
	}

	fmt.Println(slc, len(slc), cap(slc))
}
