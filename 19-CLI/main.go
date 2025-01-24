package main

import (
	"log"
	"mycli/app"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
