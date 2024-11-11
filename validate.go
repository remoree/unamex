package unamex

import (
	"errors"
	"sort"
	"strings"
)

// Validator is a function type used to define rules for validating usernames.
// Each Validator function takes a username as input and returns:
//   - A boolean indicating whether the validation passed.
//   - An error providing details if the validation fails.
//
// Example:
//
//	func MinLengthValidator(username string) (bool, error) {
//	    if len(username) < 5 {
//	        return false, errors.New("username must be at least 5 characters")
//	    }
//	    return true, nil
//	}
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

// isValid checks if the provided suggestion is valid.
// A suggestion is considered valid if it passes all validators
// and is not symmetric to the current username.
//
// If no validators are set, it applies the default validators.
//
// Returns:
//   - true if the suggestion is valid.
//   - false otherwise.
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

// isSymmetric checks if the given suggestion is the same
// as the current username. This ensures that suggestions
// are not identical to the original username.
//
// Returns:
//   - true if the suggestion is symmetric to the current username.
//   - false otherwise.
func (u *Identity) isSymmetric(suggestion string) bool {
	return u.uname == suggestion
}

// validateRange checks if the input username meets the length requirements.
// A valid username must be between 5 and 30 characters.
//
// Returns:
//   - true if the username is within the valid length range.
//   - false and an error message otherwise.
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

// validateFormat ensures that the input username follows the allowed format.
// Valid usernames can only contain letters, numbers, and one period ('.').
// Usernames cannot start or end with a period.
//
// Returns:
//   - true if the username matches the allowed format.
//   - false and an error message otherwise.
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

// validateIntegrity ensures that the username is not weak or common.
// It checks the username against a built-in blacklist of common or
// insecure usernames.
//
// Returns:
//   - true if the username is not in the blacklist.
//   - false and an error message otherwise.
func validateIntegrity(str string) (bool, error) {
	err := "username is too weak or common, please choose a different one"
	str = strings.ToLower(str)
	index := sort.SearchStrings(blacklist, str)
	if index < len(blacklist) && blacklist[index] == str {
		return false, errors.New(err)
	}

	return true, nil
}

// defaultValidator provides a default set of validators for username validation.
// These include:
//   - validateRange: Ensures the username length is valid.
//   - validateFormat: Ensures the username follows the correct format.
//   - validateIntegrity: Ensures the username is secure and not common.
//
// Returns:
//   - A slice of Validator functions representing the default validation rules.
func defaultValidator() []Validator {
	return []Validator{
		validateRange,
		validateFormat,
		validateIntegrity,
	}
}
