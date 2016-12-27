package sollozzoctl

import (
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"fmt"
	"encoding/json"
	"os"
)

var name string

var releaseCmd = &cobra.Command{
	Use:   "release [projectname] [major, minor, build]",
	Short: "Release project version",
	Long:  "Release project version",
	Run:func(cmd *cobra.Command, args[]string) {
		if len(args) != 2 {
			cmd.Help()
		} else {
			name = args[0]
			op := args[1]
			var p model.Project
			store.Get([]byte(name), &p)

			if (model.Project{}) == p {
				fmt.Print("can not found project !!")
				os.Exit(1)
			}

			switch op {
			case "major":
				p.Major += 1
			case "minor":
				p.Minor += 1
			case "build":
				p.BuildNumber += 1
			default:
				cmd.Help()
				os.Exit(1)
			}

			encode, err := json.Marshal(&p)
			if err != nil {
				panic(err)
			}

			store.Put([]byte(p.Key), encode)
			fmt.Println(p.Version())
		}
	},
}

func init() {
	cmdSollozzo.AddCommand(releaseCmd)
}