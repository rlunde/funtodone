package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// TODO: move this version to the build script (or to a constants package?)
const (
	Version = "0.0.1"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Long: `funtodone-cli is primarily intended to allow simple testing of features, by wrapping the library functions
	with a command line interface. The web app interface will be the way most real users use funtodone.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("funtodone version %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
