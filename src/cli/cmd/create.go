package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create tasks, users, or collections of tasks",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create called")
	},
}

// createUserCmd represents the create user command
var createUserCmd = &cobra.Command{
	Use:   "user",
	Short: "create a user",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create user called")
	},
}

// createTaskCmd represents the create task command
var createTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "create a task",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create task called")
	},
}

// createCollectionCmd represents the create collection command
var createCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "create a collection",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("create collection called")
	},
}

func init() {
	fmt.Println("create.go init called")
	RootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createUserCmd)
	createCmd.AddCommand(createTaskCmd)
	createCmd.AddCommand(createCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createUserCmd.Flags().String("name", "", "The user's name")
	createUserCmd.Flags().String("password", "", "The user's password")
	createUserCmd.Flags().String("email", "", "The user's email address")
	// attributes of a task or of the task's status
	createTaskCmd.Flags().String("description", "", "The task description")
	createTaskCmd.Flags().String("summary", "", "A short task summary")
	createTaskCmd.Flags().Int("level", 0, "The task level (if in a stack, the height on the stack)")
	createTaskCmd.Flags().Bool("done", false, "True if the task is done")
	createTaskCmd.Flags().Bool("started", false, "True if the task is started")
	createTaskCmd.Flags().String("due", "", "Date and time the task is due")
	createTaskCmd.Flags().String("created", "", "Date and time the task was created")
	createTaskCmd.Flags().String("modified", "", "Date and time of the last time the task was modified")
	createTaskCmd.Flags().String("completed", "", "Date and time the task was completed")
	createCollectionCmd.Flags().String("description", "", "The collection description")
	createCollectionCmd.Flags().String("summary", "", "A short summary of the collection")
	createCollectionCmd.Flags().String("type", "", "The type of collection: stack, cycle, or list")

}
