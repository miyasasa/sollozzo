package sollozzoctl

import (
	"github.com/spf13/cobra"
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
	store.Delete([]byte(args[0]))
}
