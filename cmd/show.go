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

		printSortedByDaysAway(birthdays)
	},
}

func getBirthdaysToPrint(month string) (model.Birthdays, error) {
	birthdays := make(model.Birthdays)
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

func getNextBirthday(birthday time.Time) (nextBd time.Time, age int, daysAway int) {
	var nextBdYear int

	curMonth := time.Now().Month()
	birthdayMonth := birthday.Month()
	if birthdayMonth < curMonth || (birthdayMonth == curMonth && birthday.Day() < time.Now().UTC().Day()) {
		nextBdYear = time.Now().Year() + 1
	} else {
		nextBdYear = time.Now().Year()
	}

	age = nextBdYear - birthday.Year()
	nextBd = birthday.AddDate(age, 0, 0)
	daysAway = int(nextBd.Sub(time.Now().UTC().Truncate(24*time.Hour)).Hours() / 24)

	return
}

type bdInfo struct {
	Name        string
	NextBd      time.Time
	Age         int
	DaysUntilBd int
}

func printSortedByDaysAway(birthdays model.Birthdays) {
	var temp []bdInfo

	for name, bd := range birthdays {
		nextBd, age, days := getNextBirthday(bd)
		temp = append(temp, bdInfo{
			Name:        name,
			NextBd:      nextBd,
			Age:         age,
			DaysUntilBd: days,
		})
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].DaysUntilBd < temp[j].DaysUntilBd
	})

	for _, birthday := range temp {
		if birthday.DaysUntilBd == 0 {
			fmt.Printf("-  %s's %s birthday is today! 🎂\n",
				fmt.Sprintf("\033[1m%s\033[0m", birthday.Name),
				addSuffix(birthday.Age),
			)
		} else {
			fmt.Printf("-  %s's %s birthday is %d days away, on %v\n",
				fmt.Sprintf("\033[1m%s\033[0m", birthday.Name),
				addSuffix(birthday.Age),
				birthday.DaysUntilBd,
				birthday.NextBd.Format("Jan 2, 2006"),
			)
		}
	}
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
