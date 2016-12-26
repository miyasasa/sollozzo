package sollozzoctl

import "github.com/spf13/cobra"

var currentCmd = &cobra.Command{
	Use:"current [project current version]",
	Short:"Show project current version",
	Long:"Show project current version",
	Run:runCurrentCommand,
}

func init() {
	cmdSollozzo.AddCommand(currentCmd);
}

func runCurrentCommand(cmd *cobra.Command,args []string) {

}