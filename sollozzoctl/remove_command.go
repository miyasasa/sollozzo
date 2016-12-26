package sollozzoctl

import (
	"github.com/spf13/cobra"
	"github.com/boltdb/bolt"
	"fmt"
)

var removeCmd = &cobra.Command{
	Use:"remove [remove project]",
	Short:"Remove project",
	Long:"Remove project",
	Run:runRemoveCommand,
}

func init() {
	cmdSollozzo.AddCommand(removeCmd);
}

func runRemoveCommand(cmd *cobra.Command, args []string) {

	store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("projects"))

		bucket.Delete([]byte(args[0]));

		fmt.Println(args[0], " removed")

		return nil
	})
}
