package main

import (
	"fmt"
	"time"
)

func main() {
	// counterSecs() retorna um canal de inteiros
	ch := counterSecs()
	for sec := range ch {
		if sec == 1 {
			fmt.Printf("A execução da função começou há %d segundo.\n", sec)
		} else {
			fmt.Printf("A execução da função começou há %d segundos.\n", sec)
		}
	}
	fmt.Print("Fim do programa.")
}

// Essa função conta há quantos segundos o programa está sendo executado
func counterSecs() <-chan int {
	ch := make(chan int)
	sec := 0

	go func() {
		for {
			ch <- sec
			if sec == 10 {
				close(ch) // Na main, o programa só saí do for range
				break     // porque o canal está sendo fechado aqui
			}
			time.Sleep(time.Second)
			sec++
		}
	}()
	// É retornado sempre que um valor for enviado para esse canal.
	// Nesse exemplo, ele é enviado a cada segundo.
	return ch
}
