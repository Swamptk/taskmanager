package cmd

import (
	"Task/db"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:     "do",
	Short:   "Marks a task as done.",
	Example: "Task do 1\nTask do 1 2 5",
	Long: `Marks a task as done. This will remove it from the list shown by the list command and unavailable
	for removing or doing it again manually. However, done tasks will get removed automatically from the 
	database one day after its completion.`,
	Run: func(cmd *cobra.Command, args []string) {
		allTasks, err := db.GetTasks()
		if err != nil {
			fmt.Println("Could not load the db.")
			os.Exit(1)
		}
		tasks := FilterUndone(allTasks)
		for _, i := range args {
			id, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Failed to parse the id", i)
				continue
			}
			if id < 1 || id > len(tasks) {
				fmt.Printf("The id %d is out of range.\n", id)
				continue
			}
			db.DoTask(tasks[id-1].Key)
			fmt.Printf("Task %d has been marked as done.\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
