package model

import (
	"os"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

type Project struct {
	Key         string
	Major       uint8
	Minor       uint8
	BuildNumber uint16
}

func (p *Project) Version() string {
	return fmt.Sprintf("%v.%v.%v", p.Major, p.Minor, p.BuildNumber)
}

func Display(projects []Project) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Project Name", "Version"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("*")

	for _, val := range projects {
		column := []string{val.Key, val.Version()};

		table.Append(column);
	}

	table.Render()
}
