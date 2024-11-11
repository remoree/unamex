package unamex

import (
	"errors"
	"sort"
	"strings"
)

type Validator func(string) (bool, error)

// Validate checks if the username in the Identity object is valid
// according to the given validators of:
//
//	type Validator func(string) (bool, error)
//
// If no validators are given, it uses the default ones in the Identity object.
//
// Example usage:
//
//	// Define a new validator function
//	func myValidator(username string) (bool, error) {
//	   // Implement your validator logic here
//	   if username == "" {
//	       return false, errors.New("username cannot be empty")
//	   }
//	   return true, nil
//	}
//
//	// Create a new Identity object
//	// Use the WithValidator method to replace the default
//	// validators with your own
//	u := New("myUsername").WithValidator(myValidator)
func (u *Identity) Validate(validators ...Validator) error {
	if len(validators) > 0 {
		u.validator = append(u.validator, validators...)
	}

	for _, f := range u.validator {
		if ok, err := f(u.uname); !ok {
			return err
		}
	}
	return nil
}

// WithValidator replaces the existing validators in the Identity
// object with new ones.
//
// Example usage:
//
//	// Define a new validator function
//	func myValidator(username string) (bool, error) {
//	   // Implement your validator logic here
//	   if username == "" {
//	       return false, errors.New("username cannot be empty")
//	   }
//	   return true, nil
//	}
//
//	// Create a new Identity object
//	// Use the WithValidator method to replace the default
//	// validators with your own
//	u := New("myUsername").WithValidator(myValidator)
func (u *Identity) WithValidator(validators ...Validator) *Identity {
	u.validator = validators
	return u
}

func (u *Identity) isValid(suggestion string) bool {
	if len(u.validator) <= 0 {
		u.validator = defaultValidator()
	}

	for _, f := range u.validator {
		if ok, _ := f(suggestion); !ok || u.isSymmetric(suggestion) {
			return false
		}
	}
	return true
}

func (u *Identity) isSymmetric(suggestion string) bool {
	return u.uname == suggestion
}

func validateRange(input string) (bool, error) {
	// Check if the username is empty
	if input == "" {
		return false, errors.New("username cannot be empty")
	}

	// Check if the username is too long or too short
	if len(input) < 5 || len(input) > 30 {
		return false, errors.New("username must be between 5 and 30 characters")
	}

	return true, nil
}

// Check if the username matches the format
func validateFormat(input string) (bool, error) {
	var err = "usernames can only contain letters, numbers, and period"
	const dotChar = '.'
	if input == "" || input[0] == dotChar || input[len(input)-1] == dotChar {
		return false, errors.New(err)
	}
	const limitSpecialCharacters = 1
	var countSpecialCharacters int
	var countDigit int
	for _, c := range []byte(input) {

		if c == dotChar && countSpecialCharacters < limitSpecialCharacters {
			countSpecialCharacters++
			continue
		}

		if !isLetter(c) {
			if isDigit(c) {
				countDigit++
			} else {
				return false, errors.New(err)
			}
		}
	}

	if !(countDigit != len(input)-countSpecialCharacters) {
		return false, errors.New(err)
	}

	return true, nil
}

// Check if the username is secure and not easy to guess
func validateIntegrity(str string) (bool, error) {
	err := "username is too weak or common, please choose a different one"
	str = strings.ToLower(str)
	index := sort.SearchStrings(blacklist, str)
	if index < len(blacklist) && blacklist[index] == str {
		return false, errors.New(err)
	}

	return true, nil
}

func defaultValidator() []Validator {
	return []Validator{
		validateRange,
		validateFormat,
		validateIntegrity,
	}
}
