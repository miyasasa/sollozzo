package sollozzoctl

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

func NewListCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List of your projects",
		Long:  "List of your projects",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) != 0 {
				return fmt.Errorf("\"sollozzo list\" accepts no argument(s).")
			}
			return runListCommand(store)
		},
	}

	return cmd
}
func runListCommand(store *boltdb.Store) error {

	var projects []model.Project

	store.ForEach(func(k, v []byte) error {

		var p model.Project
		json.Unmarshal(v, &p)

		projects = append(projects, p)

		return nil
	})

	if 0 == len(projects) {
		return fmt.Errorf("You do not have a project yet")
	}

	model.Display(projects);

	return nil
}
