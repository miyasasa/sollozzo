package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

func NewRemoveCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "remove [remove project]",
		Short: "Remove project",
		Long:  "Remove project",
		RunE: func(cmd *cobra.Command, args []string) error {
			//convert args to Add opts
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
