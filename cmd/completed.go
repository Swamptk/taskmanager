/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Task/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:     "completed",
	Short:   "Show a list of the tasks you have completed today.",
	Example: "Task completed",
	Long: `Show a list of the tasks you have completed today. These cannot be removed manually,
	however they will get deleted automatically a day after their completion.
	To see a list of your incomplete tasks, see list.`,
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, err := db.GetTasks()
		tasks := FilterDone(allTasks)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have not completed any task today. You should stop procrastinating 😡")
			return
		}

		for _, t := range tasks {
			fmt.Printf("✅ %s\n", t.Value.Task)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
