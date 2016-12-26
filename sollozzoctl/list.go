package sollozzoctl

import (
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

}

