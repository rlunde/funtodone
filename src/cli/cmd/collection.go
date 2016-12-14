package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// collectionCmd represents the collection command
var collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Use to add, delete, import, export, save, or load a Stack, List, Cycle, or FlashTasks",
	Long: `funtodone-cli is primarily intended to allow simple testing of features, by wrapping the library functions
	with a command line interface. The web app interface will be the way most real users use funtodone.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("collection called")
	},
}

func init() {
	RootCmd.AddCommand(collectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
