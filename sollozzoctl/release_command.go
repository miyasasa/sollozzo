package sollozzoctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

var name string
var major bool
var minor bool
var build bool

func NewReleaseCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "release <project_name> [--major, --minor, --build]",
		Short: "Release project version",
		Long:  "Release project version",
		RunE: func(cmd *cobra.Command, args []string) error {

			length := len(args)

			if length != 1 {
				return fmt.Errorf("\"sollozzo release\" accepts project name and version parameter(optional) arguments.")
			}

			name = args[0]

			return runReleaseCommand(store)
		},
	}

	cmd.Flags().BoolVarP(&major, "major", "M", false, "increment version major parameter")
	cmd.Flags().BoolVarP(&minor, "minor", "m", false, "increment version minor parameter")
	cmd.Flags().BoolVarP(&build, "build", "b", false, "increment version build parameter")

	return cmd
}
func runReleaseCommand(store *boltdb.Store) error {
	var p model.Project

	err := store.Get([]byte(name), &p)

	if err != nil {
		return fmt.Errorf("Project can not found")
	}

	if major == true {
		p.Major += 1
	}
	if minor == true {
		p.Minor += 1
	}
	if build == true {
		p.BuildNumber += 1
	}
	if major == false && minor == false && build == false {
		p.BuildNumber += 1
	}

	store.Put([]byte(p.Key), &p)

	fmt.Println(p.Version())

	return err
}
