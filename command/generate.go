package command

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/license"
)

type GenerateCommand struct {
	ID          string
	Output      string `long:"output" short:"o" ENV:"OUTPUT" description:"Output file"`
	ForceOutput bool   `long:"force-output" short:"f" ENV:"FORCE_OUTPUT" description:"Rewrite file if exists"`

	dates []string
	names []string
}

func (c *GenerateCommand) ArgsUsage() string {
	return "[license ID]"
}

func (c *GenerateCommand) Execute(context *cli.Context) {
	err := c.loadLicenseID(context)
	if err != nil {
		logrus.Fatal(err)
	}

	err = c.loadDatesAndNames(context)
	if err != nil {
		logrus.Fatal(err)
	}

	generator := c.prepareGenerator()
	err = generator.Generate()
	if err != nil {
		logrus.Fatal(err)
	}
}

func (c *GenerateCommand) loadLicenseID(context *cli.Context) error {
	args := context.Args()
	if len(args) < 1 {
		return errors.New("You must provide 'license ID' as a first command argument")
	}

	c.ID = context.Args()[0]

	return nil
}

func (c *GenerateCommand) loadDatesAndNames(context *cli.Context) error {
	c.dates = context.StringSlice("copyright-date")
	c.names = context.StringSlice("copyright-name")

	if len(c.dates) != len(c.names) {
		return errors.New("Copyright-date and copyright-name must be added in pairs")
	}

	if len(c.dates) < 1 {
		return errors.New("There must be at least one copyright-date/copyright-name pair")
	}

	return nil
}

func (c *GenerateCommand) prepareGenerator() *license.Generator {
	generator := license.NewGenerator(c.prepareLicense(), c.Output)

	if c.ForceOutput == true {
		generator.ForceOutput = true
	}

	return generator
}

func (c *GenerateCommand) prepareLicense() *license.License {
	lic := license.NewLicense(c.ID)

	counter := 0
	max := len(c.dates)
	for counter < max {
		lic.AddCopyright(license.NewCopyright(c.dates[counter], c.names[counter]))
		counter++
	}

	return lic
}

func init() {
	flags := []cli.Flag{
		cli.StringSliceFlag{
			Name:  "copyright-date, d",
			Usage: "Date of copyright owner",
		},
		cli.StringSliceFlag{
			Name:  "copyright-name, n",
			Usage: "Name of the copyright owner",
		},
	}
	RegisterCommand("generate", "Generate license", &GenerateCommand{}, flags...)
}
