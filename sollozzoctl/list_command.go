package sollozzoctl

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
)

var listCmd = &cobra.Command{
	Use:   "list [list of projects]",
	Short: "List of your projects",
	Long:  "List of your projects",
	Run:   runListCommand,
}

func init() {
	cmdSollozzo.AddCommand(listCmd)
}

func runListCommand(cmd *cobra.Command, args []string) {

	var projects []model.Project

	store.ForEach(func(k, v []byte) error {

		var p model.Project
		json.Unmarshal(v, &p)

		projects = append(projects, p)

		return nil
	})

	if 0 != len(projects) {
		for _, ps := range projects {
			fmt.Println(ps.Display())
		}
	} else {
		fmt.Println("You do not have a project yet")
	}

}
