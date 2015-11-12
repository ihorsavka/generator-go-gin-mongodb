package main

import (
	"os"

	"github.com/codegangsta/cli"

	"./server"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "<%= baseName %>"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			Value:  9000,
			Usage:  "http port",
			EnvVar: "HTTP_PORT",
		},
		cli.StringFlag{
			Name:   "mongodb",
			Usage:  "MongoDB URL",
			Value:  "localhost:27017/<%= _.slugify(baseName) %>",
			EnvVar: "MONGODB_URL",
		},
		cli.StringFlag{
			Name:  "debug",
			Usage: "Debug mode",
		},
	}
	app.Action = func(c *cli.Context) {
		port := c.Int("port")
		mongodb_url := c.String("mongodb")
		debug := c.Bool("debug")

		cfg := &server.Config{
			Port:       port,
			Debug:      debug,
			MongoDBURL: mongodb_url,
		}

		server := server.New(cfg)
		server.Start()

	}
	return app
}

func main() {
	NewApp().Run(os.Args)
}
