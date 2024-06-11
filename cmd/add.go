/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/util"
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <name> <mm/dd/yyyy>",
	Args:  cobra.ExactArgs(2),
	Short: "Add a birthday",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.StoreBirthday(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
