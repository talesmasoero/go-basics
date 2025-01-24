package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna uma aplicação de linha de comando (cli)
// pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	// Nome e uso (funcionalidade) da CLI
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e nomes de servidores na internet"
	flags := []cli.Flag{
		// Pode ter N flags de qualquer tipo
		cli.StringFlag{
			// Nome da Flag
			Name: "host",
			// Valor default da ação caso não passemos uma flag
			Value: "google.com",
		},
	}

	// Onde os comandos são configurados, é um []Command
	app.Commands = []cli.Command{
		{
			// Nome do comando usado na CLI
			Name: "buscar",
			// Funcionalidade do comando
			Usage: "Busca IPs e servidores de websites",
			// Parâmetros ([]cli.Flag)
			Flags: flags,
			// É a ação (funcionalidade) que o comando faz, sempre
			// é uma função que tenha *cli.Context como parâmetro
			Action: buscar,
		},
	}
	return app
}

func buscarIPs(c *cli.Context) {
	// Passamos o nome da flag como parâmetro
	host := c.String("host")

	// Pacote net (nativo do Go), LookupIP retorna um slice de IPs
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	// Pacote net (nativo do Go), LookupNS retorna um slice de NS
	// NS = Name Server
	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}

func buscar(c *cli.Context) {
	buscarIPs(c)
	buscarServidores(c)
}
