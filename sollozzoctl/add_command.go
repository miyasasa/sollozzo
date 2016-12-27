package sollozzoctl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
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

	var project = &model.Project{Key: args[0], Major: 1, Minor: 0, BuildNumber: 0}

	err := store.Put([]byte(args[0]), project)

	if err != nil {
		fmt.Println("Project can not created")
		os.Exit(1)
	} else {

		fmt.Println(project.Key + " created successfully")
	}
}
