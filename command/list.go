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

func (c *ListCommand) Execute(context *cli.Context) {

	logrus.Println("Available licenses:")
	for _, name := range license.ListTemplates() {
		logrus.Println(fmt.Sprintf("  %s", name))
	}
}

func init() {
	RegisterCommand("list", "List available license templates", &ListCommand{})
}
