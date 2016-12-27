package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

func NewAddCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add [Add new Project]",
		Short: "Add new Project",
		Long:  "Add new Project",
		RunE: func(cmd *cobra.Command, args []string) error {
			//convert args to Add opts
			return runAddCommand(store, args)
		},
	}

	return cmd
}
func runAddCommand(store *boltdb.Store, args []string) error {

	var project = &model.Project{Key: args[0], Major: 1, Minor: 0, BuildNumber: 0}

	err := store.Put([]byte(args[0]), project)

	if err == nil {
		fmt.Println(project.Key + " created successfully")
	}

	return err
}
