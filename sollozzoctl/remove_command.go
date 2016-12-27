package sollozzoctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var removeCmd = &cobra.Command{
	Use:   "remove [remove project]",
	Short: "Remove project",
	Long:  "Remove project",
	Run:   runRemoveCommand,
}

func init() {
	cmdSollozzo.AddCommand(removeCmd)
}

func runRemoveCommand(cmd *cobra.Command, args []string) {
	err := store.Delete([]byte(args[0]))
	if err != nil {
		fmt.Println("Project can not removed")
		os.Exit(1)
	} else {
		fmt.Println(args[0], " removed successfully")
	}
}
