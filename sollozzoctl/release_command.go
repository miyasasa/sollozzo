package sollozzoctl

import "github.com/spf13/cobra"

var releaseCmd = &cobra.Command{
	Use:"release [release version]",
	Short:"Release project version",
	Long:"Release project version",
	Run:runReleaseCommand,
}

func init() {
	cmdSollozzo.AddCommand(releaseCmd);
}

func runReleaseCommand(cmd *cobra.Command, args []string) {

}