package command

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/common"
)

type VersionCommand struct {
}

func (c *VersionCommand) ArgsUsage() string {
	return ""
}

func (c *VersionCommand) Execute(context *cli.Context) {
	info, err := common.GetVersion().ExtendedInfo()
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(info)
}

func init() {
	RegisterCommand("version", "Print version details", &VersionCommand{})
}
