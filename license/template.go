package license

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var reg = regexp.MustCompile("[A-Za-z0-9\\-\\.+]+$")

func loadTemplate(id string, useUserTemplate bool) (string, error) {
	var templateType string
	var templateBytes []byte
	var err error

	if useUserTemplate {
		templateBytes, err = loadUserTemplate(id)
		templateType = "User"
	} else {
		templateBytes, err = loadInternalTemplate(id)
		templateType = "Internal"
	}

	if err != nil {
		err := fmt.Errorf("%s template for license '%s' was not found", templateType, id)
		return "", err
	}

	return string(templateBytes), nil
}

func loadUserTemplate(id string) ([]byte, error) {
	templateFile, err := userConfigDir(fmt.Sprintf("templates/%s", id))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(templateFile)
}

func loadInternalTemplate(id string) ([]byte, error) {
	templateFile := fmt.Sprintf("templates/%s", id)

	return Asset(templateFile)
}

func ListInternalTemplates() []string {
	var templates []string
	for _, template := range AssetNames() {
		templates = append(templates, reg.FindString(template))
	}

	return templates
}

func ListUserTemplates() ([]string, error) {
	files, err := scanUserTemplatesFiles()
	if err != nil {
		return nil, err
	}

	var templates []string
	for _, f := range files {
		templates = append(templates, reg.FindString(f.Name()))
	}

	return templates, nil
}

func scanUserTemplatesFiles() ([]os.FileInfo, error) {
	userTemplatesDir, err := userConfigDir("templates")
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(userTemplatesDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, nil
	}

	return ioutil.ReadDir(userTemplatesDir)
}
