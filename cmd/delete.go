/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/util"
	"fmt"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a person's birthday",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.DeleteBirthday(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
