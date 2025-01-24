package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Println(num)
		if num == 10 {
			break
		}
		num++
	}

	var arr = [...]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	countLetters := make(map[rune]int)
	slc := []string{"João", "Pedro", "Felipe"}
	for _, nome := range slc {
		for _, char := range nome {
			countLetters[char]++
		}
		fmt.Println(nome)
	}
	for char, count := range countLetters {
		fmt.Println(string(char), count)
	}
}
