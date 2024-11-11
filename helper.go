package unamex

import (
	"math/rand"
)

const (
	// vowelsBitset         = 1065233
	asciiCaseOffset      = 0x20
	char_a          byte = 0x61
	char_A          byte = 0x41
)

// DoubleByte creates a slice of length 2, and fills it with the input byte.
// The result is a slice where both elements are the input byte.
func DoubleByte(b byte) []byte {
	const twin = 2
	var result = make([]byte, twin)
	for i := 0; i < twin && i < len(result); i++ {
		result[i] = b
	}
	return result
}

// VanishVowel iterates over each character in the input string.
// If the character is a vowel and it's not the first character,
// it removes the vowel with a 50% chance.
// If no vowel was removed during the iteration, it removes
// the last vowel in the string.
// If there are no vowels in the string or if the only vowel
// is the first character, it returns the original string.
func VanishVowel(s string) string {
	b := []byte(s)
	var lastVowelIndex int = -1

	for idx := range b {
		var c = b[idx]

		if isVowel(c) && idx != 0 {
			lastVowelIndex = idx

			if rand.Intn(2) == 1 {
				b = append(b[:idx], b[idx+1:]...)
				return string(b)
			}
		}
	}

	if lastVowelIndex > 0 {
		b = append(b[:lastVowelIndex], b[lastVowelIndex+1:]...)
	}

	return string(b)
}

// VowelTransform iterates over each character in the input string.
// If the character is a vowel, it applies the input function to the vowel
// with a 50% chance. If no vowel was transformed during the iteration,
// it applies the input function to the last vowel in the string.
// If there are no vowels in the string, it returns the original string.
//
// Example usage:
//
//		// Define a transformation function
//		func changeVowel(c byte) byte {
//	       if c == 'e' { return 'a' }; return c
//		 }
//
//		// Use the Transform function to apply the transformation
//		result := VowelTransform("hello", changeVowel)
//		fmt.Println("Transformed string:", result) // Output: hallo
func VowelTransform(s string, f func(byte) byte) string {
	b := []byte(s)
	var lastVowelIndex int = -1
	var char byte

	for i := range b {
		if isVowel(b[i]) {
			lastVowelIndex = i
			char = b[i]

			if rand.Intn(2) == 1 {
				b[i] = f(char)
				return string(b)
			}
		}
	}

	if s == string(b) {
		b[lastVowelIndex] = f(char)
	}
	return string(b)
}

// RepeatVowel iterates over each character in the input string.
// If the character is a vowel, it duplicates the vowel with a 50% chance.
// If no vowel was duplicated during the iteration,
// it duplicates the last vowel in the string.
// If there are no vowels in the string, it returns the original string.
func RepeatVowel(s string) string {
	b := []byte(s)
	var lastVowelIndex int = -1
	for idx, c := range b {
		if isVowel(c) {
			lastVowelIndex = idx
			if rand.Intn(2) == 1 {
				b = append(b[:idx+1], b[idx:]...)
				return string(b)
			}
		}
	}
	if lastVowelIndex >= 0 {
		b = append(b[:lastVowelIndex+1], b[lastVowelIndex:]...)
	}
	return string(b)
}

// Transform applies a transformation function to each character in a string.
//
// Example usage:
//
//	// Define a transformation function
//	func toUpper(c byte) byte {
//	    return byte(unicode.ToUpper(rune(c)))
//	}
//
//	// Use the Transform function to apply the transformation
//	result := Transform("hello", toUpper)
//	fmt.Println("Transformed string:", result) // Output: HELLO
func AlphabetTransform(s string, f func(byte) byte) string {
	b := []byte(s)
	for i := range b {
		b[i] = f(b[i])
	}
	// -> Convert the byte slice back to a string (this also doesn't allocate)
	// return *(*string)(unsafe.Pointer(&bytes))
	// -> Convert the byte slice back to a string (this allocates)
	return string(b)
}

// RepeatSubfix divides the input string into two halves.
// It then repeats the second half of the string and appends it to the first half.
// The result is a string where the second half is repeated once.
func RepeatSubfix(s string) string {
	var b = []byte(s)
	var m = len(b) / 2
	b = append(b[:m+1], b[m:]...)
	return string(b)
}

// RepeatPrefix repeats the first character of the string
// and appends it to the beginning of the string.
func RepeatPrefix(s string) string {
	var b = []byte(s)
	b = append(b[:1], b...)
	return string(b)
}

// RepeatSuffix repeats the last character of the string
// and appends it to the end of the string.
func RepeatSuffix(s string) string {
	var b = []byte(s)
	var suffix = b[len(b)-1:]
	b = append(b, suffix...)
	return string(b)
}

// SetPostInitialSep inserts the byte separator right after
// the first character of the input string.
func SetPostInitialSep(s string, sep byte) string {
	var b = []byte(s)
	b = append(b[:1], sep)
	b = append(b, []byte(s[1:])...)
	return string(b)
}

// SetPenultimateSep inserts the byte right before the last
// character of the input string.
func SetPenultimateSep(s string, sep byte) string {
	var b = []byte(s)
	var last = b[len(b)-1]
	b = append(append(b[:len(b)-1], sep), last)
	return string(b)
}

// SetPenultimateSepDigit appends the byte and a random digit
// to the end of the input string.
func SetPenultimateSepDigit(s string, sep byte) string {
	var b = []byte(s)
	b = append(b, sep, byte(rand.Intn(10)+'0'))
	return string(b)
}

// SetPostInitialSepDigit inserts a random digit and the byte right
// after the first character of the input string.
func SetPostInitialSepDigit(s string, sep byte) string {
	var b = []byte(s)
	digit := byte(rand.Intn(10) + '0')
	b = append(b, digit, sep)
	b = append(b[len(b)-2:], b[:len(b)-2]...)
	return string(b)
}

// SepWithRandomDigit appends the byte and a random digit from the range 0
// to nRange to the end of the input string.
func SepWithRandomDigit(s string, sep byte, nRange int) string {
	var b = []byte(s)
	b = append(b, sep)
	b = append(b, byteNumbers[rand.Intn(nRange)]...)
	return string(b)
}

// SwapTwoChars checks the length of the string, if it's less than 3,
// it returns the original string. If the last two characters are not the same,
// it swaps them. If the length of the string is even,
// it swaps the two middle characters. Otherwise, it swaps characters
// at one-third and two-thirds of the way through the string.
func SwapTwoChars(s string) string {
	var b = []byte(s)
	if len(b) < 3 {
		return s
	}

	switch {
	// If the the last two characters are not the same
	case b[len(b)-1] != b[len(b)-2]:
		// Swap the last two characters
		b[len(b)-1], b[len(b)-2] = b[len(b)-2], b[len(b)-1]
	// If the length of b is even
	case len(b)%2 == 0:
		// Swap the two middle characters
		mid := len(b) / 2
		b[mid-1], b[mid] = b[mid], b[mid-1]
	default:
		// Swap characters at one-third and two-thirds of the way
		// through the string
		oneThird := len(b) / 3
		twoThirds := 2 * oneThird
		b[oneThird], b[twoThirds] = b[twoThirds], b[oneThird]
	}

	return string(b)
}

// RepeatInitialAppendDigit repeats the first character of the string
// and appends it to the beginning of the string.
// It then appends a random digit from the range 0 to nRange to the end of the string.
func RepeatInitialAppendDigit(s string, nRange int) string {
	var b = []byte(s)
	b = append(b[:1], b...)
	// If this scheme is needed -> byte(rand.Intn(10)+'0')
	// Add ‘0’ (which is 48 in ASCII) to the random number
	// to get the correct ASCII value of the digit
	b = append(b, byteNumbers[rand.Intn(nRange)]...)
	return string(b)
}

// PrefixRandomDigit generates a random digit from the range specified by nRange.
// It then appends this digit to the end of the string.
// Finally, it moves the appended digit to the beginning of the string.
// The place of the digit in the string depends on the value of nRange.
func PrefixRandomDigit(s string, nRange int) string {
	var b = []byte(s)
	var place int = 1
	var lowerBound = 0
	if nRange > 10 && nRange <= 100 {
		place = 2
		lowerBound = 10
	} else if nRange > 100 && nRange <= 1000 {
		place = 3
		lowerBound = 100
	}

	digit := byteNumbers[rand.Intn(nRange-lowerBound)+lowerBound]
	b = append(b, digit...)
	b = append(b[len(b)-place:], b[:len(b)-place]...)
	return string(b)
}

// SuffixRandomDigit generates a random digit from the range specified by nRange.
// It then appends this digit to the end of the string.
func SuffixRandomDigit(s string, nRange int) string {
	var b = []byte(s)
	b = append(b, byteNumbers[rand.Intn(nRange)]...)
	return string(b)
}

// SetPrefixRandomDigit generates a random number between 0 and 2.
// Depending on the generated number, it calls the PrefixRandomDigit
// function with different range parameters.
func SetPrefixRandomDigit(s string) string {
	switch rand.Intn(3) {
	case 0:
		return PrefixRandomDigit(s, 1000)
	case 1:
		return PrefixRandomDigit(s, 100)
	default:
		return PrefixRandomDigit(s, 10)
	}
}

// SetSuffixRandomDigit generates a random number between 0 and 2.
// Depending on the generated number, it calls the SuffixRandomDigit
// function with different range parameters.
func SetSuffixRandomDigit(s string) string {
	switch rand.Intn(3) {
	case 0:
		return SuffixRandomDigit(s, 1000)
	case 1:
		return SuffixRandomDigit(s, 100)
	default:
		return SuffixRandomDigit(s, 10)
	}
}

// SetSepWithRandomDigit generates a random number between 0 and 2.
// Depending on the generated number, it calls the SepWithRandomDigit
// function with different range parameters.
func SetSepWithRandomDigit(s string) string {
	switch rand.Intn(3) {
	case 0:
		return SepWithRandomDigit(s, separator, 1000)
	case 1:
		return SepWithRandomDigit(s, separator, 100)
	default:
		return SepWithRandomDigit(s, separator, 10)
	}
}

func shuffleSuggestors(slice []Suggestor) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c byte) bool {
	return (c >= '0' && c <= '9')
}

func isVowel(c byte) bool {
	if !isLetter(c) {
		return false
	}

	return (vowelsBitset&(1<<uint(c-char_a)) != 0) ||
		(vowelsBitset&(1<<uint(c-char_A)) != 0)
}

func vowelSwap(c byte) byte {

	switch c {
	case 'a':
		return 97
	case 'e':
		return 'i'
	case 'i':
		return 'e'
	case 'o':
		return 'u'
	case 'u':
		return 'o'
	case 'A':
		return 65
	case 'E':
		return 'I'
	case 'I':
		return 'E'
	case 'O':
		return 'U'
	case 'U':
		return 'O'
	default:
		return '0'
	}
}

func alphabetSwap(c byte) byte {
	// Use a switch statement to handle the Replacements
	switch c {
	case 'a':
		return 'e'
	case 'e':
		return 'i'
	case 'i':
		return 'e'
	case 'o':
		return 'u'
	case 'u':
		return 'o'
	case 'c':
		return 'k'
	case 'k':
		return 'c'
	case 'g':
		return 'j'
	case 'j':
		return 'g'
	case 'p':
		return 'b'
	case 'b':
		return 'p'
	case 'f':
		return 'v'
	case 'v':
		return 'f'
	default:
		return c
	}
}
