package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

func NewRemoveCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "remove <project_name>",
		Short: "Remove project",
		Long:  "Remove project",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 1 {
				return fmt.Errorf("\"sollozzo remove\" accepts only project name argument.")
			}

			return runRemoveCommand(store, args)
		},
	}

	return cmd
}

func runRemoveCommand(store *boltdb.Store, args []string) error {

	err := store.Delete([]byte(args[0]))

	if err == nil {
		fmt.Println(args[0], " removed successfully")
	}

	return err
}
