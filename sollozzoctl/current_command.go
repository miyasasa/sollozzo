package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

func NewCurrentCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "current <project_name>",
		Short: "Show project current version",
		Long:  "Show project current version",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 1 {
				return fmt.Errorf("\"sollozzo current\" accepts only project name argument.")
			}

			return runCurrentCommand(store, args)
		},
	}

	return cmd
}
func runCurrentCommand(store *boltdb.Store, args []string) error {

	var proj model.Project

	err := store.Get([]byte(args[0]), &proj)

	if err == nil {
		fmt.Println(proj.Version())
	}
	return err
}
