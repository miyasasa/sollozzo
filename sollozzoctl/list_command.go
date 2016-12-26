package sollozzoctl

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"github.com/boltdb/bolt"
	"fmt"
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

	var projects []Project;

	store.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("projects"))
		bucket.ForEach(func(k, v []byte) error {
			var project Project;
			json.Unmarshal(v, &project);
			projects = append(projects, project);
			return nil;
		})

		return nil;
	})

	for _, ps := range projects {
		fmt.Println(ps.Display())
	}

}

