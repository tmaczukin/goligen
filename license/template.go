package license

import (
	"fmt"
	"regexp"
)

func loadTemplate(id string) (string, error) {
	templateFile := fmt.Sprintf("templates/%s", id)

	templateBytes, err := Asset(templateFile)
	if err != nil {
		err := fmt.Errorf("Templatr for license %s was not found", id)
		return "", err
	}

	return string(templateBytes), nil
}

func ListTemplates() []string {
	var templates []string

	reg := regexp.MustCompile("[A-Za-z0-9\\-\\.+]+$")

	for _, template := range AssetNames() {
		templates = append(templates, reg.FindString(template))
	}

	return templates
}
