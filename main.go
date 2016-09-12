package main

import (
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/command"
	"gitlab.com/tmaczukin/goligen/common"
	"gitlab.com/tmaczukin/goligen/helpers"
)

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

	version := common.GetVersion()
	cli.VersionPrinter = version.Printer

	app := cli.NewApp()
	helpers.AddLogLevelFlags(app)
	helpers.SetLogFormatter(app)

	app.Name = path.Base(os.Args[0])
	app.Usage = "Simple license file generator"
	app.Version = version.ShortInfo()
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

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
