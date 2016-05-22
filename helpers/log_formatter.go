package helpers

import (
	"bytes"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	COLOR_CYAN    = "\033[1;36m"
	COLOR_WHITE   = "\033[1;37m"
	COLOR_YELLOW  = "\033[0;33m"
	COLOR_BLUE    = "\033[1;34m"
	COLOR_MAGENTA = "\033[1;35m"
	COLOR_RED     = "\033[1;31m"
	COLOR_ESCAPE  = "\033[0m"
)

type GoligenFormatter struct {
	noColor    bool
	forceColor bool
}

func (gf *GoligenFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var lvlColor, lvlText string

	switch entry.Level {
	case logrus.DebugLevel:
		lvlColor = COLOR_CYAN
		lvlText = "debug"
	case logrus.InfoLevel:
		lvlColor = COLOR_WHITE
		lvlText = "info"
	case logrus.WarnLevel:
		lvlColor = COLOR_YELLOW
		lvlText = "warning"
	case logrus.ErrorLevel:
		lvlColor = COLOR_BLUE
		lvlText = "error"
	case logrus.FatalLevel:
		lvlColor = COLOR_MAGENTA
		lvlText = "fatal"
	case logrus.PanicLevel:
		lvlColor = COLOR_RED
		lvlText = "panic"
	default:
	}

	b := bytes.Buffer{}
	if gf.noColor == true && gf.forceColor != true {
		fmt.Fprintf(&b, "%s: %s", lvlText, entry.Message)
	} else {
		fmt.Fprintf(&b, "%s%s: %s%s", lvlColor, lvlText, entry.Message, COLOR_ESCAPE)
	}
	b.WriteByte('\n')

	return b.Bytes(), nil
}

func SetLogFormatter(app *cli.App) {
	newFlags := []cli.Flag{
		cli.BoolFlag{
			Name:   "no-color",
			Usage:  "Disable ANSI coloring",
			EnvVar: "NO_COLOR",
		},
		cli.BoolFlag{
			Name:   "force-color",
			Usage:  "Force ANSI coloring",
			EnvVar: "FORCE_COLOR",
		},
	}
	app.Flags = append(app.Flags, newFlags...)

	appBefore := app.Before
	app.Before = func(c *cli.Context) error {
		logrus.SetFormatter(&GoligenFormatter{
			noColor:    c.Bool("no-color"),
			forceColor: c.Bool("force-color"),
		})

		if appBefore != nil {
			return appBefore(c)
		}

		return nil
	}
}
