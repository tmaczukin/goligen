package license

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
)

func userConfigDir(suffix string) (string, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	configDir := fmt.Sprintf("%s/.goligen", homedir)
	if suffix != "" {
		configDir = fmt.Sprintf("%s/%s", configDir, suffix)
	}

	return configDir, nil
}
