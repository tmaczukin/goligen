package command

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/license"
)

type ListCommand struct {
}

func (c *ListCommand) ArgsUsage() string {
	return " "
}

func (c *ListCommand) printTemplates(templateType string, templates []string) {
	if len(templates) <= 0 {
		return
	}

	logrus.Println(fmt.Sprintf("Available %s license templates:", templateType))
	for _, name := range templates {
		logrus.Println(fmt.Sprintf("  %s", name))
	}
}

func (c *ListCommand) Execute(context *cli.Context) {
	templates := license.ListInternalTemplates()
	c.printTemplates("internal", templates)

	templates, err := license.ListUserTemplates()
	if err != nil {
		logrus.Fatal(err)
	}

	c.printTemplates("user", templates)
}

func init() {
	RegisterCommand("list", "List available license templates", &ListCommand{})
}
