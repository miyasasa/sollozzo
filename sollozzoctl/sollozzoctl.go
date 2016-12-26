package sollozzoctl

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	//TODO : change directory to ~/.sollozzo/sollozzo.db and check directory before start of the application
	path = "./sollozzo.db"

	cliName        = "sollozzo"
	cliDescription = "sollozzo is a version number generation tool for generate unique version numbers."
)

var cmdSollozzo = &cobra.Command{
	Use:   cliName,
	Short: cliDescription,

	Run: func(cCmd *cobra.Command, args []string) {
		cCmd.HelpFunc()(cCmd, args)
	},
}

func Main() {
	store := NewStore(path)

	err := store.Open()

	if err != nil {
		log.Panic(err)
	}

	defer store.Close()

	if len(os.Args) == 1 {
		cmdSollozzo.HelpFunc()(cmdSollozzo, os.Args)
		os.Exit(0)
	}

	cmdSollozzo.Execute()
}
