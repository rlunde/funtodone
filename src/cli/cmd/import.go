package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import tasks, users, or collections of tasks",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("import called")
	},
}

// importUserCmd represents the import user command
var importUserCmd = &cobra.Command{
	Use:   "user",
	Short: "import a user",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("import user called")
	},
}

// importTaskCmd represents the import task command
var importTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "import a task",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("import task called")
	},
}

// importCollectionCmd represents the import collection command
var importCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "import a collection",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("import collection called")
	},
}

func init() {
	// fmt.Println("import.go init called")
	RootCmd.AddCommand(importCmd)
	importCmd.AddCommand(importUserCmd)
	importCmd.AddCommand(importTaskCmd)
	importCmd.AddCommand(importCollectionCmd)
	importUserCmd.Flags().String("filename", "", "The file containing the user(s) to import")
	importTaskCmd.Flags().String("filename", "", "The file containing the task(s) to import")
	importCollectionCmd.Flags().String("filename", "", "The file containing the collection(s) to import")

}
