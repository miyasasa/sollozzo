package sollozzoctl

import (
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"fmt"
	"encoding/json"
)

var name string

var releaseCmd = &cobra.Command{
	Use:   "release [release version]",
	Short: "Release project version",
	Long:  "Release project version",
	Run:func(cmd *cobra.Command, args[]string) {
		if len(args) != 2 {
			cmd.Help()
		} else {
			name = args[0]
			op := args[1]
			dispatch(name, op)
		}
	},
}

func init() {
	cmdSollozzo.AddCommand(releaseCmd)
}

func dispatch(projectName string, op string) {
	var p model.Project
	store.Get([]byte(projectName), &p)

	switch op {
	case "major":
		p.Major += 1
	case "minor":
		p.Minor += 1
	case "build":
		p.BuildNumber += 1
	}

	encode, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	store.Put([]byte(p.Key), encode)
	fmt.Println(p.Version())
}