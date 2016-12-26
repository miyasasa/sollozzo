package sollozzoctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:"current [project current version]",
	Short:"Show project current version",
	Long:"Show project current version",
	Run:runCurrentCommand,
}

func init() {
	cmdSollozzo.AddCommand(currentCmd);
}

func runCurrentCommand(cmd *cobra.Command, args []string) {

	//var project = &Project{Key:args[0], Desc: "Description", Major:1, Minor: 0, BuildNumber: 0}
	var proj Project

	store.Get([]byte(args[0]), &proj)

	fmt.Println(proj.Display())
}