package sollozzoctl

import (
	"fmt"

	"encoding/json"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:"list [list of projects]",
	Short:"List of your projects",
	Long:"List of your projects",
	Run:runListCommand,
}

func init() {
	cmdSollozzo.AddCommand(listCmd)
}

func runListCommand(cmd *cobra.Command, args []string) {

	var projects []Project;

	store.forEach(func(k, v []byte) error {

		var p Project
		json.Unmarshal(v, &p)

		projects = append(projects, p)

		return nil;
	})

	for _, ps := range projects {
		fmt.Println(ps.Display())
	}

}



