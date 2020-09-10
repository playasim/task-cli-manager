package cmd

import (
	"fmt"
	"log"

	"strconv"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "remove a task from task manager",
	Run: func(cmd *cobra.Command, args []string) {
		var ids[] int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal("args error, please enter the task id.")
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
