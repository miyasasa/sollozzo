package sollozzoctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

var name string

func NewReleaseCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "release <project_name> <major, minor, build>",
		Short: "Release project version",
		Long:  "Release project version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("\"sollozzo release\" accepts project name and version parameter arguments.")
			}

			name = args[0]
			op := args[1]

			//convert args to Release  opts
			return runReleaseCommand(store, cmd,name,op)
		},
	}

	return cmd
}
func runReleaseCommand(store *boltdb.Store, cmd *cobra.Command, name string, op string) error {
		var p model.Project

		err := store.Get([]byte(name), &p)

		if err != nil {
			fmt.Print("Project can not found")
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

		store.Put([]byte(p.Key), &p)

		fmt.Println(p.Version())

	return err
}
