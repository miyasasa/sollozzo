package sollozzoctl

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yasinKIZILKAYA/sollozzo/model"
	"github.com/yasinKIZILKAYA/sollozzo/boltdb"
)

type releaseOpt struct {
	name  string
	major bool
	minor bool
	build bool
}

func NewReleaseCommand(store *boltdb.Store) *cobra.Command {

	opt := &releaseOpt{};

	cmd := &cobra.Command{
		Use:   "release <project_name> [--major, --minor, --build]",
		Short: "Release project version",
		Long:  "Release project version",
		RunE: func(cmd *cobra.Command, args []string) error {

			length := len(args)

			if length != 1 {
				return fmt.Errorf("\"sollozzo release\" accepts project name and version parameter(optional) arguments.")
			}

			opt.name = args[0]

			return runReleaseCommand(store, opt)
		},
	}

	cmd.Flags().BoolVarP(&opt.major, "major", "M", false, "increment version major parameter")
	cmd.Flags().BoolVarP(&opt.minor, "minor", "m", false, "increment version minor parameter")
	cmd.Flags().BoolVarP(&opt.build, "build", "b", false, "increment version build parameter")

	return cmd
}
func runReleaseCommand(store *boltdb.Store, opt *releaseOpt) error {
	var p model.Project

	err := store.Get([]byte(opt.name), &p)

	if err != nil {
		return fmt.Errorf("Project can not available")
	}

	if opt.major == true {
		p.Major += 1
	}
	if opt.minor == true {
		p.Minor += 1
	}
	if opt.build == true {
		p.BuildNumber += 1
	}
	if opt.major == false && opt.minor == false && opt.build == false {
		p.BuildNumber += 1
	}

	store.Put([]byte(p.Key), &p)

	fmt.Println(p.Version())

	return err
}
