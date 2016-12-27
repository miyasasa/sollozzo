package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
)

var currentCmd = &cobra.Command{
	Use:   "current [project current version]",
	Short: "Show project current version",
	Long:  "Show project current version",
	Run:   runCurrentCommand,
}

func init() {
	cmdSollozzo.AddCommand(currentCmd)
}

func runCurrentCommand(cmd *cobra.Command, args []string) {
	var proj model.Project

	err := store.Get([]byte(args[0]), &proj)

	if err != nil {
		fmt.Print("Project can not found")
	} else {
		fmt.Println(proj.Display())
	}
}
