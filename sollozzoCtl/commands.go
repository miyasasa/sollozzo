package sollozoCtl

import "github.com/spf13/cobra"

var rootCmd=&cobra.Command{}

func Main()  {
	rootCmd.Execute();
}