/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

var (
	task_input string
)

type taskcsv struct{
	task string `csv:"task"`
	
}
// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(task_input)
		writecsv()
	},
}

func writecsv(){
	
	file, err := os.Create("task-list1.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	inptasks := []*taskcsv{
		{ task_input},
	}

	if err := gocsv.MarshalFile(inptasks, file);err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().StringVarP(&task_input, "create", "c", "", "Enter your task")
}
