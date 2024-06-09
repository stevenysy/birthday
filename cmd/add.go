/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/model"
	"birthday/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a birthday",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				os.Exit(1)
			}
			return
		}

		if len(args) > 2 {
			fmt.Println("Too many arguments\n")
			usage()
			return
		}

		birthday, err := model.NewBirthday(args[0], args[1])
		if err != nil {
			fmt.Println(err)
		}
		util.StoreBirthday(birthday)
	},
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  birthday add [name] [date]")
}

func help(cmd *cobra.Command, _ []string) {
	fmt.Println(cmd.Short + "\n")

	usage()
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.SetHelpFunc(help)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
