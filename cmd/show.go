/*
Copyright © 2024 Steven Yi stevenjxhc@gmail.com
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

var month string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show birthdays",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		birthdays, err := getBirthdaysToPrint(month)
		if err != nil {
			fmt.Println(err)
			return
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

func getBirthdaysToPrint(month string) ([]model.Birthday, error) {
	var birthdays []model.Birthday
	var err error

	if month == "" {
		birthdays, err = util.ReadAllBirthdays()
		if err != nil {
			return nil, err
		}
	} else {
		birthdays, err = util.ReadBirthdays(month)
		if err != nil {
			return nil, err
		}
	}

	if len(birthdays) == 0 {
		return nil, fmt.Errorf("no birthdays found in %s 😢", month)
	}
	return birthdays, nil
}

func getNextBirthday(birthday model.Birthday) (nextBd time.Time, age int, daysAway int) {
	var nextBdYear int

	curMonth := time.Now().Month()
	birthdayMonth := birthday.Date.Month()
	if birthdayMonth < curMonth || (birthdayMonth == curMonth && birthday.Date.Day() < time.Now().UTC().Day()) {
		nextBdYear = time.Now().Year() + 1
	} else {
		nextBdYear = time.Now().Year()
	}

	age = nextBdYear - birthday.Date.Year()
	nextBd = birthday.Date.AddDate(age, 0, 0)
	daysAway = int(nextBd.Sub(time.Now().UTC().Truncate(24*time.Hour)).Hours() / 24)

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

	showCmd.Flags().StringVarP(&month, "month", "m", "", "display birthdays of one month only")
}
