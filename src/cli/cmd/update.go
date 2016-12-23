package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update tasks, users, or collections of tasks",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update called")
	},
}

// updateUserCmd represents the update user command
var updateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "update a user",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update user called")
	},
}

// updateTaskCmd represents the update task command
var updateTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "update a task",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update task called")
	},
}

// updateCollectionCmd represents the update collection command
var updateCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "update a collection",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("update collection called")
	},
}

func init() {
	fmt.Println("update.go init called")
	RootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateUserCmd)
	updateCmd.AddCommand(updateTaskCmd)
	updateCmd.AddCommand(updateCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
