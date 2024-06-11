package util

import (
	"birthday/errors"
	"birthday/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"time"
)

func StoreBirthday(birthday *model.Birthday) error {
	f, err := openBirthdayFile()

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Read birthdays from file and unmarshal to birthdays slice if not empty
	var birthdays []model.Birthday
	fileBytes, err := os.ReadFile(getBirthdayFileDir())
	if err != nil {
		return fmt.Errorf("error reading birthdays.json: %v", err)
	}
	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &birthdays)
		if err != nil {
			return fmt.Errorf("error unmarshaling birthdays.json: %v", err)
		}
	}

	// If the person already has a birthday set, we remove it to update it with the new birthday
	for i, b := range birthdays {
		if b.Name == birthday.Name {
			birthdays = append(birthdays[:i], birthdays[i+1:]...)
		}
	}
	birthdays = append(birthdays, *birthday)

	// Write updated birthdays to file
	err = writeToBirthdayFile(birthdays, f, false)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBirthday(name string) error {
	f, err := openBirthdayFile()

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var birthdays []model.Birthday

	fileBytes, err := os.ReadFile(getBirthdayFileDir())
	if err != nil {
		return fmt.Errorf("error reading birthdays.json: %v", err)
	}

	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &birthdays)
		if err != nil {
			return fmt.Errorf("error unmarshaling birthdays.json: %v", err)
		}

		if deleteIx := findBirthdayIndex(name, birthdays); deleteIx != -1 {
			birthdays = append(birthdays[:deleteIx], birthdays[deleteIx+1:]...)
		}

		err = writeToBirthdayFile(birthdays, f, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadAllBirthdays() ([]model.Birthday, error) {
	b, err := os.ReadFile(getBirthdayFileDir())
	if err != nil {
		return nil, errors.ErrNoBirthdays
	}

	var birthdays []model.Birthday
	err = json.Unmarshal(b, &birthdays)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling birthdays.json: %v", err)
	}

	if len(birthdays) == 0 {
		return nil, errors.ErrNoBirthdays
	}

	return birthdays, nil
}

func ReadBirthdays(month string) ([]model.Birthday, error) {
	allBds, err := ReadAllBirthdays()
	if err != nil {
		return nil, err
	}

	var birthdaysOfMonth []model.Birthday
	filter, err := time.Parse("January", month)
	if err != nil {
		return nil, fmt.Errorf("incorrect month format, please enter a full month name")
	}

	for _, bd := range allBds {
		if bd.Date.Month() == filter.Month() {
			birthdaysOfMonth = append(birthdaysOfMonth, bd)
		}
	}

	return birthdaysOfMonth, nil
}

func openBirthdayFile() (*os.File, error) {
	f, err := os.OpenFile(getBirthdayFileDir(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening birthdays.json: %v", err)
	}
	return f, nil
}

func writeToBirthdayFile(birthdays []model.Birthday, f *os.File, overwrite bool) error {
	b, err := json.Marshal(birthdays)
	if err != nil {
		return fmt.Errorf("error marshaling birthdays.json: %v", err)
	}

	if overwrite {
		f, err = os.Create("birthdays.json")
		if err != nil {
			return fmt.Errorf("error truncating birthdays.json: %v", err)
		}
	}

	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("error writing to birthdays.json: %v", err)
	}

	return nil
}

func findBirthdayIndex(name string, birthdays []model.Birthday) int {
	for i, b := range birthdays {
		if b.Name == name {
			return i
		}
	}
	return -1
}

func getBirthdayFileDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(usr.HomeDir, ".birthdays.json")
}
