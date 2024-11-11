package example

import (
	"errors"

	"github.com/remoree/unamex"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Initialize a printer for formatted output in English
var fmtPrinter = message.NewPrinter(language.English)
var out = fmtPrinter.Println

// Check validates a username and generates suggestions if invalid or unavailable.
// Input:
//   - username: The username to validate.
//
// Returns:
//   - A slice of suggested usernames if validation fails.
//   - An error if validation fails or the username is unavailable.
func Check(username string) (suggestions []string, err error) {
	// Ensure the username is not empty
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	// Create a new Identity instance with the given username
	u := unamex.New(username)

	// Validate the username with custom availability rules
	if err := u.Validate(withAvailabilityCheck); err != nil {
		// If the error is due to availability, generate suggestions
		if errors.Is(err, errAvailability) {
			capacity := 5 // Number of suggestions to generate
			return u.Suggest(capacity), err
		}

		// Return any other validation error
		return nil, err
	}

	// Print success message if username is valid
	out("Validation passed. Username is valid.")
	return nil, nil
}

// withAvailabilityCheck is a custom validation rule that checks if
// a username is available.
// Input:
//   - s: The username to check.
//
// Returns:
//   - false and errAvailability if the username is unavailable.
//   - true if the username is available.
func withAvailabilityCheck(s string) (bool, error) {
	// Mark specific usernames as unavailable
	if s == "moree" { // Example reserved username
		return false, errAvailability
	}

	// Indicate that the username is available
	return true, nil
}

// errAvailability is a custom error returned when a username is unavailable.
var errAvailability = errors.New(
	"this username is unavailable, please choose a different one")
