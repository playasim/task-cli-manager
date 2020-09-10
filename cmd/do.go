package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "remove a task from task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
