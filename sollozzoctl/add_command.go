package sollozzoctl

import (
	"fmt"
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/boltdb/bolt"
)

var addCmd = &cobra.Command{
	Use:   "add [Add new Project]",
	Short: "Add new Project",
	Long:  "Add new Project",
	Run:   runAddCommand,
}

func init() {
	cmdSollozzo.AddCommand(addCmd)
}

func runAddCommand(cmd *cobra.Command, args []string) {
	//var version = Version{1, 0, 0}
	var project = &Project{Key:args[0], Desc: "Description", Major:1, Minor: 0, BuildNumber: 0}

	store.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("projects"))

		encoded, err := json.Marshal(&project)

		if err != nil {
			return err
		}
		bucket.Put([]byte(project.Key), encoded)
		return nil;
	})

	fmt.Println(project.Key + " created successfully")
}
