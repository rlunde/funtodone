package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save tasks, users, or collections of tasks",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
	real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("save called")
	},
}

// saveUserCmd represents the save user command
var saveUserCmd = &cobra.Command{
	Use:   "user",
	Short: "save a user",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("save user called")
	},
}

// saveTaskCmd represents the save task command
var saveTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "save a task",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("save task called")
	},
}

// saveCollectionCmd represents the save collection command
var saveCollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "save a collection",
	Long: `funtodone-cli is primarily a wrapper around library functions. Most
    real users will only use it via the web app interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("save collection called")
	},
}

func init() {
	fmt.Println("save.go init called")
	RootCmd.AddCommand(saveCmd)
	saveCmd.AddCommand(saveUserCmd)
	saveCmd.AddCommand(saveTaskCmd)
	saveCmd.AddCommand(saveCollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
