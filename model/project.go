package model

import "fmt"

type Project struct {
	Key         string
	Desc        string
	Major       uint8
	Minor       uint8
	BuildNumber uint16
}

func (p *Project) Display() string {
	return fmt.Sprintf("%v %v.%v.%v", p.Key, p.Major, p.Minor, p.BuildNumber)
}

func (p *Project) Version() string {
	return fmt.Sprintf("%v.%v.%v", p.Major, p.Minor, p.BuildNumber)
}
