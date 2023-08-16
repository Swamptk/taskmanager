/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Task/db"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Removes a task from the list.",
	Example: "Task rm 1\nTask rm 1 2 4",
	Long: `Removes a task from the database using the key shown using the list command. 
	This command is only available for uncompleted tasts, since the completed ones are removed
	from the db automatically a day after they get completed. Accepts an arbitrary number of 
	indexes to remove.`,
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
			db.RmTask(tasks[id-1].Key)
			fmt.Printf("Task %d has been removed.\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
