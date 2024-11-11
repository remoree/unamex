package example

import (
	"errors"

	"github.com/remoree/unamex"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var fmt = message.NewPrinter(language.English)
var out = fmt.Println

func Check(username string) (suggestions []string, err error) {
	u := unamex.New(username)

	if err := u.Validate(withAvailabilityCheck); err != nil {
		if errors.Is(err, errAvailability) {
			capacity := 5
			return u.Suggest(capacity), err
		}

		return nil, err
	}

	print("pass...")

	return nil, nil
}

func withAvailabilityCheck(s string) (bool, error) {
	if s == "moree" {
		return false, errAvailability
	}

	return true, nil
}

var errAvailability = errors.New(
	"this username is unavailable, please choose a different one")
