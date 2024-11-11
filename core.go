// Package unamex offers flexible and efficient tools for username validation
// and suggestion generation. It ensures usernames adhere to configurable
// rules, such as length and format, while preventing weak or common choices.
//
// The package also provides highly customizable suggestion strategies,
// allowing developers to generate alternative usernames based on specific
// requirements or patterns.
//
// Example usage:
//
//	u := unamex.New("exampleUser")
//	if err := u.Validate(); err != nil {
//		fmt.Println("Invalid username:", err)
//		fmt.Println("Suggestions:", u.Suggest(5))
//	}
package unamex

// Identity represents the core structure for username validation
// and suggestion. It provides mechanisms to validate usernames
// against custom rules and generate alternative suggestions.
type Identity struct {
	// uname is the username being validated or used for suggestions.
	uname string

	// validator is a slice of Validator functions used to validate
	// the username. These functions define the rules for a valid
	// username, such as length, format, etc.
	validator []Validator

	// suggestor is a slice of Suggestor functions used to generate
	// alternative username suggestions. These functions define
	// the strategies for creating suggestions, such as adding prefixes
	// or modifying vowels.
	suggestor []Suggestor
}

// Suggestor is a function type used to define strategies
// for generating username suggestions. Each Suggestor function
// takes a username as input and returns a modified or alternative username.
//
// Example:
//
//	func AddSuffix(username string) string {
//	    return username + "_123"
//	}
type Suggestor func(s string) string

// New creates a new Identity instance with default validator and suggestors.
// An optional username can be provided as an argument.
// If a username is provided, it will be set as the uname of the new Identity instance.
// If no username is provided, the 'uname' field will be set to "default".
// You can also set or change the uname later using the 'On' method.
func New(username ...string) *Identity {
	u := &Identity{
		uname:     "default",
		validator: defaultValidator(),
		suggestor: defaultSuggestors(),
	}

	if len(username) > 0 {
		u.uname = username[0]
	}

	return u
}

// On sets the uname field of a Identity instance
// to the provided username string and returns
// the updated Identity instance.
func (u *Identity) On(username string) *Identity {
	u.uname = username
	return u
}

// WithSuggestor replaces the existing suggestors in the Identity
// object with new ones or adds new suggestors of:
//
//	type Suggestor func(s string) string
//
// Example usage:
//
//	// Define a new suggestor function
//	func mySuggestor(username string) string {
//	   // Implement your suggestor logic here
//	   return "suggestedUsername"
//	}
//	// Create a new Identity object.
//	// Use the WithSuggestor method to replace
//	// the default suggestors with your own
//	u := New("myUsername").WithSuggestor(mySuggestor)
func (u *Identity) WithSuggestor(suggestors ...Suggestor) *Identity {
	u.suggestor = suggestors
	return u
}

// Suggest first checks if there are any suggestors passed as input.
// If there are, it appends them to the suggestor field of the Identity struct.
// It then shuffles the suggestors in the suggestor field.
// The method creates a slice of strings with a length equal to the capacity input.
// It then iterates over the suggestors up to the capacity input.
// The capacity will be adjusted to match the number of available suggestors
// if it initially exceeds that number.
// For each iteration, it calls the suggestor with the 'uname' field of the Identity
// struct as input. If the suggestion is valid according to all validators
// in the validator field of the Identity struct, it adds the suggestion
// to the suggestions slice. The method returns the suggestions slice.
func (u *Identity) Suggest(capacity int, suggestors ...Suggestor) []string {
	if len(suggestors) > 0 {
		u.suggestor = append(u.suggestor, suggestors...)
	}

	if capacity > len(u.suggestor) {
		capacity = len(u.suggestor)
	}

	shuffleSuggestors(u.suggestor)

	suggestions := make([]string, 0, capacity)

	seen := make(map[string]bool)

	var suggestor Suggestor

	var suggestion string

	for i := 0; i < capacity; i++ {
		suggestor = u.suggestor[i]
		suggestion = suggestor(u.uname)

		if !u.isValid(suggestion) {
			continue
		}

		if !seen[suggestion] {
			suggestions = append(suggestions, suggestion)
			seen[suggestion] = true
		}
	}

	return suggestions
}

// defaultSuggestors returns a slice of default Suggestor functions.
// Each Suggestor implements a unique strategy to generate alternative
// usernames by modifying the input username in various ways.
func defaultSuggestors() []Suggestor {
	var suggestors = []Suggestor{
		func(s string) string { return SetPrefixRandomDigit(s) },
		func(s string) string { return SetSuffixRandomDigit(s) },
		func(s string) string { return SetSepWithRandomDigit(s) },

		func(s string) string { return SetPenultimateSep(s, separator) },
		func(s string) string { return SetPostInitialSep(s, separator) },

		func(s string) string { return SetPenultimateSepDigit(s, separator) },
		func(s string) string { return SetPostInitialSepDigit(s, separator) },

		func(s string) string { return SwapTwoChars(s) },

		func(s string) string { return RepeatPrefix(s) },
		func(s string) string { return RepeatSuffix(s) },
		func(s string) string { return RepeatSubfix(s) },
		func(s string) string { return RepeatVowel(s) },
		func(s string) string { return RepeatInitialAppendDigit(s, 10) },

		func(s string) string { return AlphabetTransform(s, alphabetSwap) },
		func(s string) string { return VowelTransform(s, vowelSwap) },

		func(s string) string { return VanishVowel(s) },
	}

	return suggestors
}
