/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/model"
	"birthday/util"
	"fmt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <name> <birthday>",
	Args:  cobra.ExactArgs(2),
	Short: "Add a birthday",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		birthday, err := model.NewBirthday(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}

		err = util.StoreBirthday(birthday)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
