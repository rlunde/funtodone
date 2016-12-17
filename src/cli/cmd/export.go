package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export tasks, users, or collections of tasks",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("export called")
	},
}

// exportUserCmd represents the export user command
var exportUserCmd = &cobra.Command{
	Use:   "user",
	Short: "export a user",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("export user called")
	},
}

// exportTaskCmd represents the export task command
var exportTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "export a task",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("export task called")
	},
}

// exportCollectionCmd represents the export collection command
var exportCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "export a collection",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("export collection called")
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
	exportCmd.AddCommand(exportUserCmd)
	exportCmd.AddCommand(exportTaskCmd)
	exportCmd.AddCommand(exportCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
