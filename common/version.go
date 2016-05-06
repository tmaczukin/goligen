package common

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
	"time"
)

var extendedInfoTemplate = `Version:      {{.Version}}
Git revision: {{.Revision}}
GO version:   {{.GoVersion}}
Built:        {{.Built}}
OS/Arch:      {{.Os}}/{{.Arch}}`

type Version struct {
	Version   string
	Revision  string
	Built     time.Time
	GoVersion string
	Os        string
	Arch      string
}

func (v *Version) SetValues(version, revision, built string) (err error) {
	v.Version = version
	v.Revision = revision

	if built == "now" {
		v.Built = time.Now()
	} else {
		v.Built, err = time.Parse(time.RFC3339, built)
	}

	return err
}

func (v *Version) ShortInfo() string {
	return fmt.Sprintf("%s (%s)", v.Version, v.Revision)
}

func (v *Version) ExtendedInfo() (string, error) {
	tmpl, err := template.New("version-info").Parse(extendedInfoTemplate)
	if err != nil {
		return "", err
	}

	result := &bytes.Buffer{}
	err = tmpl.Execute(result, v)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

var instance *Version

func GetVersion() *Version {
	if instance == nil {
		instance = &Version{
			GoVersion: runtime.Version(),
			Os:        runtime.GOOS,
			Arch:      runtime.GOARCH,
		}
	}

	return instance
}

func init() {
	GetVersion()
}
