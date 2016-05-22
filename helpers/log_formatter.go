package helpers

import (
	"bytes"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var (
	COLORCYAN    = "\033[1;36m"
	COLORWHITE   = "\033[1;37m"
	COLORYELLOW  = "\033[0;33m"
	COLORBLUE    = "\033[1;34m"
	COLORMAGENTA = "\033[1;35m"
	COLORRED     = "\033[1;31m"
	COLORESCAPE  = "\033[0m"
)

type GoligenFormatter struct {
	noColor    bool
	forceColor bool
}

func (gf *GoligenFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	lvlColor, lvlText := getLvlColorAndText(entry.Level)

	b := bytes.Buffer{}
	if gf.noColor == true && gf.forceColor != true {
		fmt.Fprintf(&b, "%s: %s", lvlText, entry.Message)
	} else {
		fmt.Fprintf(&b, "%s%s: %s%s", lvlColor, lvlText, entry.Message, COLORESCAPE)
	}
	b.WriteByte('\n')

	return b.Bytes(), nil
}

func getLvlColorAndText(level logrus.Level) (lvlColor, lvlText string) {
	text := map[logrus.Level]string{
		logrus.DebugLevel: "debug",
		logrus.InfoLevel:  "info",
		logrus.WarnLevel:  "warning",
		logrus.ErrorLevel: "error",
		logrus.FatalLevel: "fatal",
		logrus.PanicLevel: "panic",
	}

	color := map[logrus.Level]string{
		logrus.DebugLevel: COLORCYAN,
		logrus.InfoLevel:  COLORWHITE,
		logrus.WarnLevel:  COLORYELLOW,
		logrus.ErrorLevel: COLORBLUE,
		logrus.FatalLevel: COLORMAGENTA,
		logrus.PanicLevel: COLORRED,
	}

	if text[level] != "" {
		lvlText = text[level]
	}

	if color[level] != "" {
		lvlColor = color[level]
	}

	return
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
