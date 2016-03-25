package main

import (
	"fmt"
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/command"
	"gitlab.com/tmaczukin/goligen/helpers"
)

var NAME = path.Base(os.Args[0])
var VERSION = "dev"
var REVISION = "HEAD"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			_, ok := r.(*logrus.Entry)
			if ok {
				os.Exit(1)
			}
			panic(r)
		}
	}()

	app := cli.NewApp()
	app.Name = NAME
	app.Usage = "Simple license file generator"
	app.Version = fmt.Sprintf("%s (%s)", VERSION, REVISION)
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tomasz Maczukin",
			Email: "tomasz@maczukin.pl",
		},
	}

	app.Commands = command.GetCommands()
	app.CommandNotFound = func(context *cli.Context, command string) {
		logrus.Fatalln("Command not found:", command)
	}

	helpers.AddLogLevelFlags(app)

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
