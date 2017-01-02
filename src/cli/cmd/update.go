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
	updateUserCmd.Flags().String("id", "", "The id of the user to update")
	updateTaskCmd.Flags().String("id", "", "The id of the task to update")
	updateCollectionCmd.Flags().String("id", "", "The id of the collection to update")
	updateUserCmd.Flags().String("name", "", "The user's name")
	updateUserCmd.Flags().String("password", "", "The user's password")
	updateUserCmd.Flags().String("email", "", "The user's email address")
	// attributes of a task or of the task's status
	updateTaskCmd.Flags().String("description", "", "The task description")
	updateTaskCmd.Flags().String("summary", "", "A short task summary")
	updateTaskCmd.Flags().Int("level", 0, "The task level (if in a stack, the height on the stack)")
	updateTaskCmd.Flags().Bool("done", false, "True if the task is done")
	updateTaskCmd.Flags().Bool("started", false, "True if the task is started")
	updateTaskCmd.Flags().String("due", "", "Date and time the task is due")
	updateTaskCmd.Flags().String("created", "", "Date and time the task was created")
	updateTaskCmd.Flags().String("modified", "", "Date and time of the last time the task was modified")
	updateTaskCmd.Flags().String("completed", "", "Date and time the task was completed")
	updateCollectionCmd.Flags().String("description", "", "The collection description")
	updateCollectionCmd.Flags().String("summary", "", "A short summary of the collection")
	updateCollectionCmd.Flags().String("type", "", "The type of collection: stack, cycle, or list")
}
