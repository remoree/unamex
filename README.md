# Unamex
### A Flexible Username Validation and Suggestion Library for Go


![Build Status](https://github.com/remoree/unamex/actions/workflows/ci.yml/badge.svg) ![Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen) ![Performance](https://img.shields.io/badge/Performance-High-brightgreen) ![Go Report Card](https://goreportcard.com/badge/github.com/remoree/unamex)


**Unamex** is a Go library for robust username validation and generation, ensuring uniqueness and compliance with configurable rules.


---

### Features

- **Validation**:
  - Checks username length (e.g., between 5 and 30 characters).
  - Ensures a proper format (allowing letters, numbers, and one period).
  - Detects weak or common usernames using a built-in blacklist.

- **Suggestions**:
  - Generates alternative usernames using built-in or custom suggestion algorithms.
  - Supports transformations like adding prefixes, suffixes, or modifying vowels.

- **Customizable**:
  - Define your own validation rules (`Validator` functions).
  - Add custom suggestion algorithms (`Suggestor` functions).

- **Designed for high performance and thread-safe concurrent use**
    - **Unamex** is rigorously tested and benchmarked. Below are some highlights:

        + 100% Test Coverage: Every feature, edge case, and helper function is covered.

        + Edge Case Validations: Includes validation for various username scenarios, such as blacklist matching, invalid formats, and secure usernames.

        + Parallel Tests: Many tests are designed to run in parallel, ensuring thread safety.



### Installation

Ensure your project is using Go modules (`go mod init <module-name>` if not already initialized).

Install the package using `go get`:

```bash
go get github.com/remoree/unamex
```



### Basic Usage

Here’s an example of how to use **Unamex** to validate and suggest usernames.

#### Example: Validate a Username

```go
package main

import (
	"fmt"
	"log"
	"coda/unamex"
)

func main() {
	username := "test.user"
	u := unamex.New(username)

	// Validate the username with the default rules
	err := u.Validate()
	if err != nil {
		log.Printf("Invalid username: %v\n", err)

		// Generate suggestions for the invalid username
		suggestions := u.Suggest(5) // Generate up to 5 suggestions
		fmt.Println("Suggestions:", suggestions)
		return
	}

	fmt.Println("Username is valid!")
}
```



#### Example: Custom Validation and Suggestions

You can add custom rules for validation and suggestions.

```go
package main

import (
	"errors"
	"fmt"
	"coda/unamex"
)

func main() {
	username := "example.user"
	u := unamex.New(username)

	// Add a custom validation rule
	u.WithValidator(func(input string) (bool, error) {
		if len(input) > 20 {
			return false, errors.New("username is too long")
		}
		return true, nil
	})

	// Add a custom suggestion algorithm
	u.WithSuggestor(func(s string) string {
		return s + "_123"
	})

	// Validate the username
	err := u.Validate()
	if err != nil {
		fmt.Println("Invalid username:", err)
		fmt.Println("Suggestions:", u.Suggest(3)) // Generate 3 suggestions
		return
	}

	fmt.Println("Username is valid!")
}
```



### Key Concepts

#### 1. **Validation**
Validation ensures that usernames meet specific requirements. By default, the following rules are applied:

- **Length**: Between 5 and 30 characters.

- **Format**: Must consist of letters, numbers, and up to one period `.`.

- **Integrity**: Must not be a weak or common username (based on a built-in blacklist).

You can define custom rules using `Validator` functions:

```go
type Validator func(string) (bool, error)
```



#### 2. **Suggestions**
When a username is invalid or unavailable, suggestions are generated using built-in or custom algorithms. Built-in suggestors include:

- Adding random digits as prefixes or suffixes.

- Transforming vowels.

- Repeating characters or segments.

You can define custom algorithms using `Suggestor` functions:

```go
type Suggestor func(string) string
```



#### 3. **Extensibility**
You can extend the library with custom rules and algorithms:

- **Custom Validators**: Replace or add rules with `WithValidator`.

- **Custom Suggestors**: Replace or add algorithms with `WithSuggestor`.



---



### API Reference

#### Creating an `Identity` Instance
```go
func New(username ...string) *Identity
```
Creates a new `Identity` instance for managing username validation and suggestion.



#### Setting or Updating a Username
```go
func (u *Identity) On(username string) *Identity
```
Sets or updates the username for validation.



#### Adding Custom Validators
```go
func (u *Identity) WithValidator(validators ...Validator) *Identity
```
Replaces the default validators with custom ones.



#### Adding Custom Suggestors
```go
func (u *Identity) WithSuggestor(suggestors ...Suggestor) *Identity
```
Replaces the default suggestors with custom ones.



#### Validating a Username
```go
func (u *Identity) Validate(validators ...Validator) error
```
Validates the username using default or provided rules.



#### Generating Suggestions
```go
func (u *Identity) Suggest(capacity int, suggestors ...Suggestor) []string
```
Generates up to `capacity` suggestions using default or provided algorithms.

---

### Benchmarks

Unamex is optimized for performance, with benchmarks demonstrating its efficiency for sequential and parallel operations:

Here are some benchmark results for **Unamex**:

| Benchmark                     | Iterations  | Time (ns/op) | Bytes Allocated | Allocations |
|-------------------------------|-------------|--------------|-----------------|-------------|
| BenchmarkS_New-16             | 7,129,189   | 166.6        | 216             | 3           |
| BenchmarkP_New-16             | 11,230,796  | 121.9        | 216             | 3           |
| BenchmarkS_Suggest-16         | 1,869,153   | 638.9        | 29              | 2           |
| BenchmarkP_Suggest-16         | 2,377,364   | 474.1        | 32              | 2           |
| BenchmarkS_dSuggestors-16     | 17,493,631  | 68.78        | 128             | 1           |
| BenchmarkP_dSuggestors-16     | 19,854,486  | 63.68        | 128             | 1           |
| BenchmarkS_DoubleByte-16      | 282,782,588 | 4.223        | 0               | 0           |
| BenchmarkP_DoubleByte-16      | 69,160,147  | 18.42        | 0               | 0           |
| BenchmarkS_VanishVowel-16     | 12,045,190  | 97.74        | 8               | 1           |
| BenchmarkP_VanishVowel-16     | 9,127,910   | 134.5        | 8               | 1           |



### Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.



### License

This project is licensed under the MIT License. See the `LICENSE` file for details.


### Acknowledgments

- Built with inspiration from modern username validation and suggestion practices.
- Special thanks to contributors and the Go community for their support.


