package license

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/Sirupsen/logrus"
)

type Generator struct {
	ForceOutput bool

	license *License
	output  string
	tmpl    *template.Template
}

func (g *Generator) Generate() error {
	var err error

	err = g.prepareTemplate()
	if err != nil {
		return err
	}

	logrus.Infoln("Generating license:", g.license.ID)
	if g.output != "" {
		return g.generateToFile()
	}

	return g.generateToStdout()
}

func (g *Generator) prepareTemplate() error {
	var err error
	var templateText string

	templateText, err = loadTemplate(g.license.ID)
	if err != nil {
		return err
	}

	g.tmpl, err = template.New("license-template").Parse(templateText)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateToFile() error {
	var err error

	logrus.Infoln("Generating to file:", g.output)

	_, err = os.Stat(g.output)
	if err == nil && g.ForceOutput != true {
		return fmt.Errorf("File '%s' already exists. Use -f to force rewrite", g.output)
	}

	result := &bytes.Buffer{}
	err = g.tmpl.Execute(result, g.license)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(g.output, result.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateToStdout() error {
	logrus.Infoln("Generating to standard output")

	return g.tmpl.Execute(os.Stdout, g.license)
}

func NewGenerator(license *License, output string) *Generator {
	return &Generator{
		license: license,
		output:  output,
	}
}
