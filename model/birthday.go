package model

import (
	"fmt"
	"time"
)

type Birthday struct {
	Name string    `json:"name"`
	Date time.Time `json:"birthday"`
}

func NewBirthday(name string, birthday string) (*Birthday, error) {
	t, err := time.Parse("01/02/2006", birthday)
	if err != nil {
		t, err = time.Parse("1/2/2006", birthday)
		if err != nil {
			return nil, fmt.Errorf("unrecognized date format in birthday: %w", err)
		}
	}

	return &Birthday{
		Name: name,
		Date: t,
	}, nil
}
