package license

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	DefaultLicenseID  string       `toml:"default_license_id"`
	DefaultCopyrights []*Copyright `toml:"default_copyrights"`
}

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

func LoadUserConfig() (*Config, error) {
	configFile, err := userConfigDir("config.toml")
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(configFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return nil, nil
	}

	var config Config
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
