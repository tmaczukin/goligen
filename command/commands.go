package command

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/ayufan/golang-cli-helpers"
)

var commands []cli.Command

type Command interface {
	ArgsUsage() string
	Execute(c *cli.Context)
}

func RegisterSimpleCommand(command cli.Command) {
	logrus.Debugln("Registering command:", command.Name)
	commands = append(commands, command)
}

func RegisterCommand(name, usage string, data Command, flags ...cli.Flag) {
	RegisterSimpleCommand(cli.Command{
		Name:      name,
		Usage:     usage,
		Action:    data.Execute,
		Flags:     append(flags, clihelpers.GetFlagsFromStruct(data)...),
		ArgsUsage: data.ArgsUsage(),
	})
}

func GetCommands() []cli.Command {
	return commands
}
