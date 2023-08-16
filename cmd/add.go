package cmd

import (
	"Task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a task to your task list.",
	Example: "Task add Walk the dog",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		_, err := db.CreateTask(taskName)
		if err != nil {
			fmt.Println("There was an error creating your task", taskName)
			fmt.Println(err)
			return
		}
		fmt.Printf("Added %q to your task list.", taskName)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
