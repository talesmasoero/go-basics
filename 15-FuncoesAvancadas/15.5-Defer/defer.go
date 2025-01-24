package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Print("World!")
	defer fmt.Println("Teste")
	fmt.Print("Hello, ")

	file, err := os.Open("D:/Ceub/3-semestre/ltp-2/pratica-1.docx")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
