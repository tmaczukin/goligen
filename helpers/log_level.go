package helpers

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func AddLogLevelFlags(app *cli.App) {
	newFlags := []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Set debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:   "log-level",
			Value:  "info",
			Usage:  "Set log level (options: debug, info, warn, error, fatal, panic)",
			EnvVar: "LOG_LEVEL",
		},
	}
	app.Flags = append(app.Flags, newFlags...)

	appBefore := app.Before
	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stderr)

		level, err := logrus.ParseLevel(c.String("log-level"))
		if err != nil {
			logrus.Fatalf(err.Error())
		}

		logrus.SetLevel(level)
		if !c.IsSet("log-level") && c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if appBefore != nil {
			return appBefore(c)
		}

		return nil
	}
}
