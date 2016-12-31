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
	Short: "export a user record, by serializing it to stdout",
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
	Short: "export a task, by serializing it to stdout",
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
	Short: "export a collection, by serializing it to stdout",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("export collection called")
	},
}

func init() {
	fmt.Println("export.go init called")
	RootCmd.AddCommand(exportCmd)
	exportCmd.AddCommand(exportUserCmd)
	exportCmd.AddCommand(exportTaskCmd)
	exportCmd.AddCommand(exportCollectionCmd)
	exportUserCmd.Flags().String("id", "", "The id of the user to export")
	exportTaskCmd.Flags().String("id", "", "The id of the task to export")
	exportCollectionCmd.Flags().String("id", "", "The id of the collection to export")
}
