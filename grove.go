package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "grove"
	app.Version = Version
	app.Usage = ""
	app.Author = "handlename"
	app.Email = "nagata<at>handlena.me"
	app.Action = doMain
	app.Run(os.Args)
}

func doMain(c *cli.Context) {

}
