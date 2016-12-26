package sollozoCtl

import (
	"github.com/spf13/cobra"
	"os/exec"
	"log"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:"add [Add new Project]",
	Short:"Add new Project",
	Long:"Add new Project",
	Run:func(cmd *cobra.Command, args []string) {

		id, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}

		//project := domain.Project{string(id), args[0], time.Now(), 1.0, 0.0, 0.0, 0.0, }
		//message := service.AddProject(&project);

		//fmt.Println("Project " + args[0] + " added ", message);
	},
}
