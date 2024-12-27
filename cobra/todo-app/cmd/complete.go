/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todoapp/internal/todo"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete task by id",
	Long: `coplete task by id
	eg:
	complete 1`,
	Run: completeTodoCmd,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTodoCmd(cmd *cobra.Command, args []string) {
	fmt.Println("complete called")
	id, _ := strconv.Atoi(args[0])
	todo.Complete(id)
}
