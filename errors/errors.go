package errors

import (
	"fmt"
)

var (
	ErrNoBirthdays = fmt.Errorf("no birthdays set yet, add one now! 🎉")
)
