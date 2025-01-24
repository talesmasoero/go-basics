package main

import (
	"fmt"
	"time"
)

func main() {
	exemplo1()
	fmt.Println("==================")
	exemplo2()
}

// Função trabalhadores (worker) que processa as tarefas/trabalhos
// tarefas é um canal exclusivamente de leitura (podemos ler os dados do canal)
// resultados é um canal exclusivamente de escrita (podemos enviar dados para o canal)
func trabalhadores(trabalhador int, tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		fmt.Printf("O trabalhador %d começou a tarefa %d\n", trabalhador, tarefa)
		time.Sleep(time.Second * 5) // Tempo de execução da tarefa
		fmt.Printf("O trabalhador %d terminou a tarefa %d\n", trabalhador, tarefa)

		resultados <- tarefa
	}
}

func exemplo1() {
	const numTarefas = 10

	tarefas := make(chan int, numTarefas)
	resultados := make(chan int, numTarefas)

	// Inicializa a pool com 5 workers/trabalhadores (goorutines)
	for trabalhador := 1; trabalhador <= 4; trabalhador++ {
		go trabalhadores(trabalhador, tarefas, resultados)
	}

	// Envia 20 tarefas para a pool de workers
	for tarefa := 1; tarefa <= numTarefas; tarefa++ {
		tarefas <- tarefa
	}
	close(tarefas)

	// Apenas lê os resultados das tarefas processadas
	for r := 1; r <= numTarefas; r++ {
		<-resultados
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// <-chan => canal de, apenas, leitura de dados (recebe dados do canal)
// chan<- => canal de, apenas, escrita de dados (envia dados para o canal)
func worker(tarefas <-chan int, resultados chan<- int) {
	// tarefa é um número inteiro
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func exemplo2() {
	tarefas := make(chan int, 50)
	resultados := make(chan int, 50)

	for i := 0; i < 100; i++ {
		go worker(tarefas, resultados)
	}

	for i := 0; i < 50; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 50; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
