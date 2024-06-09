package util

import (
	"birthday/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func StoreBirthday(birthday *model.Birthday) {
	f, err := os.OpenFile("birthdays.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("error opening birthdays.json: %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var birthdays []model.Birthday

	fileBytes, err := os.ReadFile("birthdays.json")
	if err != nil {
		fmt.Printf("error reading birthdays.json: %v", err)
	}

	if len(fileBytes) > 0 {
		err = json.Unmarshal(fileBytes, &birthdays)
		if err != nil {
			fmt.Printf("error unmarshaling birthdays.json: %v", err)
		}
	}

	birthdays = append(birthdays, *birthday)
	b, err := json.Marshal(birthdays)
	_, err = f.Write(b)
	if err != nil {
		fmt.Printf("error writing to birthdays.json: %v", err)
	}
}
