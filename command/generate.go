package command

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"gitlab.com/tmaczukin/goligen/license"
)

type GenerateCommand struct {
	ID              string
	Output          string `long:"output" short:"o" ENV:"OUTPUT" description:"Output file"`
	ForceOutput     bool   `long:"force-output" short:"f" ENV:"FORCE_OUTPUT" description:"Rewrite file if exists"`
	UseUserTemplate bool   `long:"use-user-template" short:"u" ENV:"USE_USER_TEMPLATE" description:"Use user template instead of internal"`

	config     *license.Config
	copyrights []*license.Copyright
}

func (c *GenerateCommand) ArgsUsage() string {
	return "[license ID]"
}

func (c *GenerateCommand) Execute(context *cli.Context) {
	var err error
	c.config, err = license.LoadUserConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	err = c.loadLicenseID(context)
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
		if c.config != nil && c.config.DefaultLicenseID != "" {
			logrus.Infoln("Using default License ID from user's configuration file")
			c.ID = c.config.DefaultLicenseID

			return nil
		}
		return errors.New("You must provide 'license ID' as a first command argument")
	}

	c.ID = context.Args()[0]

	return nil
}

func (c *GenerateCommand) loadDatesAndNames(context *cli.Context) error {
	dates := context.StringSlice("copyright-date")
	names := context.StringSlice("copyright-name")

	if len(dates) != len(names) {
		return errors.New("Copyright-date and copyright-name must be added in pairs")
	}

	if len(dates) < 1 {
		if c.config != nil && len(c.config.DefaultCopyrights) > 0 {
			logrus.Infoln("Using default Copyrights from user's configuration file")
			c.copyrights = c.config.DefaultCopyrights

			return nil
		}
		return errors.New("There must be at least one copyright-date/copyright-name pair")
	}

	prepareCOpyrights(dates, names)

	return nil
}

func (c *GenerateCommand) prepareCopyrights(dates, names []string) {
	counter := 0
	max := len(dates)
	for counter < max {
		c.copyrights = append(c.copyrights, license.NewCopyright(dates[counter], names[counter]))
		counter++
	}
}

func (c *GenerateCommand) prepareGenerator() *license.Generator {
	generator := license.NewGenerator(c.prepareLicense(), c.Output)

	if c.ForceOutput == true {
		generator.ForceOutput = true
	}

	return generator
}

func (c *GenerateCommand) prepareLicense() *license.License {
	lic := license.NewLicense(c.ID, c.UseUserTemplate)

	for _, copyright := range c.copyrights {
		lic.AddCopyright(copyright)
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
