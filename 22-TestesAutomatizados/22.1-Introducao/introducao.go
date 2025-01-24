package main

import (
	"basics/22-TestesAutomatizados/22.1-Introducao/endereco"
	"fmt"
)

func main() {
	tests := []string{"RUA DOS BOBOS", "casa 4", "AP 707"}
	for _, test := range tests {
		endereco := endereco.TipoDeEndereco(test)
		fmt.Println(endereco)
	}
}
