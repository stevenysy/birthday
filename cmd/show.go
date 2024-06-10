/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"birthday/model"
	"birthday/util"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show birthdays",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		birthdays, err := util.ReadAllBirthdays()
		if err != nil {
			fmt.Println(err)
		}

		sortByDaysAway(birthdays)

		for _, birthday := range birthdays {
			nextBd, age, daysAway := getNextBirthday(birthday)

			if daysAway == 0 {
				fmt.Printf("-  %s's %s birthday is today! 🎂\n",
					fmt.Sprintf("\033[1m%s\033[0m", birthday.Name),
					addSuffix(age),
				)
			} else {
				fmt.Printf("-  %s's %s birthday is %d days away, on %v\n",
					fmt.Sprintf("\033[1m%s\033[0m", birthday.Name),
					addSuffix(age),
					daysAway,
					nextBd.Format("Jan 2, 2006"),
				)
			}
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
	daysAway = int(nextBd.Sub(time.Now().UTC()).Hours() / 24)

	return
}

func sortByDaysAway(birthdays []model.Birthday) []model.Birthday {
	sort.Slice(birthdays, func(i, j int) bool {
		_, _, days1 := getNextBirthday(birthdays[i])
		_, _, days2 := getNextBirthday(birthdays[j])
		return days1 < days2
	})

	return birthdays
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
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
