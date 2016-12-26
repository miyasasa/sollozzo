package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"encoding/json"
)

var addCmd = &cobra.Command{
	Use:   "add [Add new Project]",
	Short: "Add new Project",
	Long:  "Add new Project",
	Run:   runAddCommand,
}

func init() {
	cmdSollozzo.AddCommand(addCmd)
}

func runAddCommand(cmd *cobra.Command, args []string) {
	//var version = Version{1, 0, 0}
	var project = &Project{Key:args[0], Desc: "Description", Major:1, Minor: 0, BuildNumber: 0}

	content, _ := json.Marshal(project)

	store.Put([]byte(args[0]), content)

	fmt.Println(project.Key + " created successfully")
}
