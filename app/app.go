package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	app.Commands = []cli.Command{ //vai ser um slice de comandos que a aplicação vai entender e conseguir executar
		{
			Name:  "ip", //comando que busca os ips
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{ //flag que vai receber os valores para o comando funcionar
				cli.StringFlag{ //parametros da flag
					Name:  "host",
					Value: "amazon.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//NET
	ips, erro := net.LookupHost(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		log.Println(ip)
	}

}
