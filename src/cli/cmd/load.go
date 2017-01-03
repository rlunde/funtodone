package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load tasks, users, or collections of tasks from the database into memory",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("load called")
	},
}

// loadUserCmd represents the load user command
var loadUserCmd = &cobra.Command{
	Use:   "user",
	Short: "load a user from the database into memory",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("load user called")
	},
}

// loadTaskCmd represents the load task command
var loadTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "load a task from the database into memory",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("load task called")
	},
}

// loadCollectionCmd represents the load collection command
var loadCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "load a collection from the database into memory",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("load collection called")
	},
}

func init() {
	// fmt.Println("load.go init called")
	RootCmd.AddCommand(loadCmd)
	loadCmd.AddCommand(loadUserCmd)
	loadCmd.AddCommand(loadTaskCmd)
	loadCmd.AddCommand(loadCollectionCmd)
	loadUserCmd.Flags().String("name", "", "The user's name")
	loadTaskCmd.Flags().String("summary", "", "A short task summary")
	loadCollectionCmd.Flags().String("summary", "", "A short summary of the collection")
	loadUserCmd.Flags().String("id", "", "The user id")
	loadTaskCmd.Flags().String("id", "", "The task id")
	loadCollectionCmd.Flags().String("id", "", "The collection id")

}
