/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/model"
	"birthday/util"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get birthdays",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		birthdays, err := util.ReadAllBirthdays()
		if err != nil {
			fmt.Println(err)
		}

		for _, birthday := range birthdays {
			nextBd, age, daysAway := getNextBirthday(birthday)

			fmt.Printf("-  %s's %s birthday is %d days away, on %d/%d/%d\n",
				birthday.Name,
				addSuffix(age),
				daysAway,
				nextBd.Month(),
				nextBd.Day(),
				nextBd.Year(),
			)
		}
	},
}

func getNextBirthday(birthday model.Birthday) (nextBd time.Time, age int, daysAway int) {
	var nextBdYear int

	curMonth := time.Now().Month()
	birthdayMonth := birthday.Birthday.Month()
	if birthdayMonth < curMonth || (birthdayMonth == curMonth && birthday.Birthday.Day() < time.Now().Day()) {
		nextBdYear = time.Now().Year() + 1
	} else {
		nextBdYear = time.Now().Year()
	}

	age = nextBdYear - birthday.Birthday.Year()
	nextBd = birthday.Birthday.AddDate(age, 0, 0)
	daysAway = int(nextBd.Sub(time.Now()).Hours() / 24)

	return
}

func addSuffix(num int) string {
	switch num % 10 {
	case 1:
		return strconv.Itoa(num) + "st"
	case 2:
		return strconv.Itoa(num) + "nd"
	case 3:
		return strconv.Itoa(num) + "rd"
	default:
		return strconv.Itoa(num) + "th"
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
