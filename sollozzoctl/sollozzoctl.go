package sollozzoctl

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

const (
	cliName = "sollozzo"
	cliDescription = "sollozzo is a version number generation tool for generate unique version numbers."
)

var cli *SollozzoCli

type SollozzoCli struct {
	cmd *cobra.Command
}

func NewSollozzoCli(s *boltdb.Store) *SollozzoCli {
	cli = &SollozzoCli{
		cmd: &cobra.Command{
			Use:   cliName,
			Short: cliDescription,

			Run: func(cCmd *cobra.Command, args []string) {
				cCmd.HelpFunc()(cCmd, args)
			},
		},
	}

	registerCommands(cli.cmd, s)

	return cli
}

func registerCommands(cmd *cobra.Command, store *boltdb.Store) {
	cmd.AddCommand(
		NewAddCommand(store),
		NewCurrentCommand(store),
		NewListCommand(store),
		NewReleaseCommand(store),
		NewRemoveCommand(store),
	)
}

func (cli *SollozzoCli) registerCommand(cmd *cobra.Command) {

}

func (cli *SollozzoCli) Execute() {
	if len(os.Args) == 1 {
		cli.cmd.HelpFunc()(cli.cmd, os.Args)
		os.Exit(0)
	}

	cli.cmd.Execute()
}
