package license

type Copyright struct {
	Date string
	Name string
}

func NewCopyright(date string, name string) *Copyright {
	return &Copyright{
		Date: date,
		Name: name,
	}
}

type License struct {
	ID              string
	UseUserTemplate bool
	Copyrights      []*Copyright
}

func (l *License) AddCopyright(c *Copyright) {
	l.Copyrights = append(l.Copyrights, c)
}

func NewLicense(id string, useUserTemplate bool) *License {
	return &License{
		ID:              id,
		UseUserTemplate: useUserTemplate,
	}
}
