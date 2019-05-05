package main

import (
	"github.com/urfave/cli"
	_ "go-crawler/config"
	"go-crawler/crawler"
	"go-crawler/database"
	"os"
)

func main() {

	database.MongoDB = database.NewConn()

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name: "get-list",
			Action: func(c *cli.Context) (err error) {
				err = crawler.GetListNews()
				return
			},
		},
		{
			Name: "get-content",
			Action: func(c *cli.Context) (err error) {
				err = crawler.GetNewsContent()
				return
			},
		},
	}

	app.Run(os.Args)
}