package app

import "github.com/urfave/cli"

// Gerar return app for command line
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.usage = "Find IPs and names of internet server"

	return app
}
