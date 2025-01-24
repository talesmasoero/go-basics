package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Pacote main")
	auxiliar.PrintAuxiliar()

	err := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(err)
}
