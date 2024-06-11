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

func StoreBirthday(name string, date string) error {
	f, err := openBirthdayFile()

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Read birthdays from file and unmarshal to birthdays slice if not empty
	birthdays := make(model.Birthdays)
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

	t, err := parseBirthday(date)
	if err != nil {
		return err
	}
	birthdays[name] = *t

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

	birthdays := make(model.Birthdays)

	fileBytes, err := os.ReadFile(getBirthdayFileDir())
	if err != nil {
		return fmt.Errorf("error reading birthdays.json: %v", err)
	}

	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &birthdays)
		if err != nil {
			return fmt.Errorf("error unmarshaling birthdays.json: %v", err)
		}

		delete(birthdays, name)

		err = writeToBirthdayFile(birthdays, f, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadAllBirthdays() (model.Birthdays, error) {
	b, err := os.ReadFile(getBirthdayFileDir())
	if err != nil {
		return nil, errors.ErrNoBirthdays
	}

	birthdays := make(model.Birthdays)
	err = json.Unmarshal(b, &birthdays)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling birthdays.json: %v", err)
	}

	if len(birthdays) == 0 {
		return nil, errors.ErrNoBirthdays
	}

	return birthdays, nil
}

func ReadBirthdays(month string) (model.Birthdays, error) {
	allBds, err := ReadAllBirthdays()
	if err != nil {
		return nil, err
	}

	birthdaysOfMonth := make(model.Birthdays)
	filter, err := time.Parse("January", month)
	if err != nil {
		return nil, fmt.Errorf("incorrect month format, please enter a full month name")
	}

	for key, value := range allBds {
		if value.Month() == filter.Month() {
			birthdaysOfMonth[key] = value
		}
	}

	return birthdaysOfMonth, nil
}

func parseBirthday(birthday string) (*time.Time, error) {
	t, err := time.Parse("01/02/2006", birthday)
	if err != nil {
		t, err = time.Parse("1/2/2006", birthday)
		if err != nil {
			return nil, fmt.Errorf("unrecognized date format in birthday: %w", err)
		}
	}
	return &t, nil
}

func openBirthdayFile() (*os.File, error) {
	f, err := os.OpenFile(getBirthdayFileDir(), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening birthdays.json: %v", err)
	}
	return f, nil
}

func writeToBirthdayFile(birthdays model.Birthdays, f *os.File, overwrite bool) error {
	b, err := json.Marshal(birthdays)
	if err != nil {
		return fmt.Errorf("error marshaling birthdays.json: %v", err)
	}

	if overwrite {
		f, err = os.Create(getBirthdayFileDir())
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

func getBirthdayFileDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(usr.HomeDir, ".birthdays.json")
}
