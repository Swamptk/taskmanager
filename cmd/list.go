package cmd

import (
	"Task/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show a list of all your tasks.",
	Example: "Task list",
	Long: `Show a list of all your incompleted tasks, with the indexes you can use to access them with
	other commands (like do, rm). To show a list of your completed tasks, see completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, err := db.GetTasks()
		tasks := FilterUndone(allTasks)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("There are no pending tasks. You should take a break 🏖️")
			return
		}
		for i, t := range tasks {
			fmt.Printf("%d - %s\n", i+1, t.Value.Task)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func FilterUndone(tasks []db.Task) []db.Task {
	ret := []db.Task{}
	for _, t := range tasks {
		if !t.Value.Done {
			ret = append(ret, t)
		}
	}
	return ret
}

func FilterDone(tasks []db.Task) []db.Task {
	ret := []db.Task{}
	for _, t := range tasks {
		if t.Value.Done {
			ret = append(ret, t)
		}
	}
	return ret
}
