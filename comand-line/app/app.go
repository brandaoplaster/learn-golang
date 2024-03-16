package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplication CLI"
	app.Usage = "Search IPs and names of servers"

	return app
}
