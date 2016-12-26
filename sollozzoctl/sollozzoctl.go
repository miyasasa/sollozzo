package sollozzoctl

import (
	"log"

	"os"
	"os/user"

	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

const (
	config = ".sollozzo"
	db     = "sollozzo.db"

	cliName        = "sollozzo"
	cliDescription = "sollozzo is a version number generation tool for generate unique version numbers."
)

var (
	store       *boltdb.Store
	cmdSollozzo = &cobra.Command{
		Use:   cliName,
		Short: cliDescription,

		Run: func(cCmd *cobra.Command, args []string) {
			cCmd.HelpFunc()(cCmd, args)
		},
	}
)

func path() string {
	current, _ := user.Current()

	return current.HomeDir + string(filepath.Separator) + config
}

func dbPath() string {
	return path() + string(filepath.Separator) + db
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func Main() {
	exist, err := exists(path())

	if err != nil {
		log.Panic(err)
	}

	if !exist {
		os.Mkdir(path(), 0755)
	}

	store = boltdb.NewStore(dbPath())

	err = store.Open()

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
