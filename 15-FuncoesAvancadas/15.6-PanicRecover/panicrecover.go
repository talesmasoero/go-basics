package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Iniciando a execução...")
	time.Sleep(time.Second / 2)

	executando()

	time.Sleep(time.Second / 2)
	fmt.Println("Finalizando a execução...")
}

func executando() {
	defer recuperarExecucao()
	fmt.Println("Durante a execução...")
	time.Sleep(time.Second / 2)
	panic("Algo deu errado, o programa entrou em panic")
	// qualquer coisa abaixo desse panic seria ignorado
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Valor/Erro capturado no panic:", r)
	}
}
