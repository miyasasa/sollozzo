package sollozzoctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"regexp"
	"strconv"
	"strings"
)

func NewAddCommand(store *boltdb.Store) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "add <project_name> [version as major.minor.build]",
		Short: "Add new Project",
		Long:  "Add new Project",
		RunE: func(cmd *cobra.Command, args []string) error {

			length := len(args)

			if length > 2 || length == 0 {
				return fmt.Errorf("\"sollozzo add\" accepts only project name and project version(optional) arguments.")
			}
			return runAddCommand(store, args)
		},
	}

	return cmd
}
func runAddCommand(store *boltdb.Store, args []string) error {

	var project = &model.Project{}

	if len(args) == 2 {

		reg, err := regexp.MatchString("^(\\d)(.\\d){2}$", args[1])

		if err != nil || reg == false {
			return fmt.Errorf("Project can not created. Invalid version number, try as 1.2.3")
		}
		vParameters := strings.Split(args[1], ".")

		major, _ := strconv.ParseUint(vParameters[0], 10, 8)
		minor, _ := strconv.ParseUint(vParameters[1], 10, 8)
		build, _ := strconv.ParseUint(vParameters[2], 10, 16)

		project = &model.Project{Key: args[0], Major: uint8(major), Minor: uint8(minor), BuildNumber: uint16(build)}

	} else {

		project = &model.Project{Key: args[0], Major: 1, Minor: 0, BuildNumber: 0}
	}

	var existProject = model.Project{}
	err := store.Get([]byte(project.Key), &existProject)

	if err == nil {
		return fmt.Errorf("Project exist. Please remove availabile project firstly")
	}

	err = store.Put([]byte(args[0]), project)

	if err == nil {
		fmt.Println(project.Key + " created successfully")
	}

	return err
}
