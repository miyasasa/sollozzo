package sollozzoctl

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
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

	store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("projects"))

		v := bucket.Get([]byte(args[0]))

		json.Unmarshal(v, &proj)

		return nil;
	})

	fmt.Println(proj.Display())
}