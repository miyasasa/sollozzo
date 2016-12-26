package sollozzoctl

import (
	"github.com/spf13/cobra"
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

	//project := domain.Project{string(id), args[0], time.Now(), 1.0, 0.0, 0.0, 0.0, }
	//message := service.AddProject(&project);

	//fmt.Println("Project " + args[0] + " added ", message);
}
