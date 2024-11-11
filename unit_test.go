package unamex

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var mockValidator = func(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username cannot be empty")
	}
	return true, nil
}

func mockSuggestor(username string) string {
	return "suggestedUsername"
}

// A mock function to check if a username is already taken by another user
func withUsernameAvailablilityCheck(username string) (bool, error) {
	err := "this username is unavailable"
	switch username {
	case "morree":
		return false, errors.New(err)
	case "morie":
		return false, errors.New(err)
	case "more":
		return false, errors.New(err)
	case "1moree":
		return false, errors.New(err)
	case "moreee":
		return false, errors.New(err)
	case "moree":
		return false, errors.New(err)
	case "murii":
		return false, errors.New(err)
	default:
		return true, errors.New(err)
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("DefaultUname", func(t *testing.T) {
		t.Parallel()
		u := New()
		require.Equal(t, "default", u.uname)
	})

	t.Run("SetUname", func(t *testing.T) {
		t.Parallel()
		u := New("testuser")
		require.Equal(t, "testuser", u.uname)
	})

	t.Run("DefaultValidator", func(t *testing.T) {
		t.Parallel()
		u := New()
		require.NotNil(t, u.validator)
	})

	t.Run("DefaultSuggestor", func(t *testing.T) {
		t.Parallel()
		u := New()
		require.NotNil(t, u.suggestor)
	})

}

func TestOn(t *testing.T) {
	username := "newuser"
	u := New()
	for i := 0; i < 5; i++ {
		u.On(username + strconv.Itoa(i))
		if i <= 3 {
			require.NotEqual(t, username, u.uname)
		}
		if i > 3 {
			require.Equal(t, username+strconv.Itoa(i), u.uname)
		}
	}
}

func TestWithValidator(t *testing.T) {
	t.Parallel()
	u := New()

	t.Run("OverwriteNil", func(t *testing.T) {
		t.Parallel()
		u.WithValidator()
		require.Nil(t, u.validator)
	})
	t.Run("CreateNew", func(t *testing.T) {
		t.Parallel()
		u.WithValidator(mockValidator)
		require.NotNil(t, u.validator)
	})
}

func TestWithSuggestor(t *testing.T) {
	t.Parallel()
	u := New()

	t.Run("OverwriteNil", func(t *testing.T) {
		// t.Parallel()
		u.WithSuggestor()
		require.Nil(t, u.suggestor)
	})
	t.Run("CreateNew", func(t *testing.T) {
		// t.Parallel()
		u.WithSuggestor(mockSuggestor)
		require.NotNil(t, u.suggestor)
	})
}

func TestValidate(t *testing.T) {
	t.Parallel()

	var unameTestCases = []struct {
		username string
		state    bool
	}{
		// Valid Cases
		{username: "sarah.adams", state: true},
		{username: "sarah.123", state: true},
		// Hack blacklist cases
		{username: "root", state: false},
		{username: "abcde", state: false},
		{username: "delete", state: false},
		// Break format rule cases
		{username: "sarah.", state: false},
		{username: ".sarah", state: false},
		{username: ".sarah.", state: false},
		{username: "sarah.adams.", state: false},
		{username: "sarah..adams", state: false},
		{username: "123456", state: false},
		{username: "11111111", state: false},
		{username: "root@test.com", state: false},
		{username: "............", state: false},
		{username: "", state: false},
		// Break range cases
		{username: "sarah1sarah2sarah3adams4.adams5", state: false},
		{username: "sar", state: false},
		{username: "abc", state: false},
		{username: "x.1", state: false},
	}
	u := New()

	t.Run("AppendExtraValidator", func(t *testing.T) {
		currentLengthValidatorBefore := len(u.validator)

		u.On(username).Validate(mockValidator)
		addOneValidator := 1
		currentLengthValidatorAfter := len(u.validator)

		require.Equal(t,
			currentLengthValidatorBefore+addOneValidator,
			currentLengthValidatorAfter)
	})

	t.Run("CheckDefaultRules", func(t *testing.T) {
		for _, v := range unameTestCases {
			err := u.On(v.username).Validate()
			if !v.state {
				require.Error(t, err)
				continue
			}
			require.NoError(t, err)
		}
	})

	t.Run("isValid", func(t *testing.T) {
		// Overwrite the default validator with nil
		u.WithValidator()
		// Expected validator length is zero
		zeroLength := 0
		require.Len(t, u.validator, zeroLength)
		// Set a username
		u.On("moree")
		// Set a mock suggested username
		u.isValid("moree")
		// Expected validator length is > 0 with default validators
		require.Greater(t, len(u.validator), zeroLength)
	})

}

func TestSuggest(t *testing.T) {
	t.Run("AppendExtraSuggestor", func(t *testing.T) {
		u := New()
		suggestors := []Suggestor{mockSuggestor, mockSuggestor}
		currentLengthSuggestors := len(u.suggestor)
		addedSuggestorLength := len(suggestors)
		u.Suggest(5, suggestors...)

		require.Equal(t,
			currentLengthSuggestors+addedSuggestorLength, len(u.suggestor))
	})
	t.Run("CheckUsernameVariants", func(t *testing.T) {
		u := New("moree")
		u.Validate(withUsernameAvailablilityCheck)
		suggestions := u.Suggest(100)
		for _, v := range suggestions {
			if v != "" {
				if u.isSymmetric(v) {
					t.Errorf("Expected variant, got %s", v)
				}
			}
		}
	})
}

func TestHelpers(t *testing.T) {
	t.Run("DoubleByte", func(t *testing.T) {
		cc := DoubleByte('c')
		require.Equal(t, "cc", string(cc))
	})

	t.Run("VanishVowel", func(t *testing.T) {
		// Since VanishVowel runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			vv := VanishVowel("every")
			require.Equal(t, "evry", vv)
		}
	})

	t.Run("VowelTransform", func(t *testing.T) {
		// Since VowelTransform runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			vt := VowelTransform("world", vowelSwap)
			require.Equal(t, "wurld", vt)
		}
	})

	t.Run("RepeatVowel", func(t *testing.T) {
		// Since RepeatVowel runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			rv := RepeatVowel("world")
			require.Equal(t, "woorld", rv)
		}
	})

	t.Run("AlphabetTransform", func(t *testing.T) {
		at := AlphabetTransform("hello", alphabetSwap)
		require.Equal(t, "hillu", at)
	})

	t.Run("AlphabetTransform", func(t *testing.T) {
		at := AlphabetTransform("hello", alphabetSwap)
		require.Equal(t, "hillu", at)
	})

	t.Run("RepeatSubfix", func(t *testing.T) {
		rs := RepeatSubfix("username")
		require.Equal(t, "usernname", rs)
	})

	t.Run("RepeatPrefix", func(t *testing.T) {
		rp := RepeatPrefix("user")
		require.Equal(t, "uuser", rp)
	})

	t.Run("RepeatSuffix", func(t *testing.T) {
		rs := RepeatSuffix("user")
		require.Equal(t, "userr", rs)
	})

	t.Run("SetPostInitialSep", func(t *testing.T) {
		sps := SetPostInitialSep("user", '.')
		require.Equal(t, "u.ser", sps)
	})

	t.Run("SetPenultimateSep", func(t *testing.T) {
		sps := SetPenultimateSep("user", '.')
		require.Equal(t, "use.r", sps)
	})

	t.Run("SetPenultimateSepDigit", func(t *testing.T) {
		var sep byte = '.'
		spsd := SetPenultimateSepDigit("user", sep)
		str := strings.Split(spsd, string(sep))
		integer, err := strconv.Atoi(str[len(str)-1])
		require.NoError(t, err)
		require.Equal(t,
			"user"+string(sep)+strconv.Itoa(integer), spsd)
	})

	t.Run("SetPostInitialSepDigit", func(t *testing.T) {
		var sep byte = '.'
		spsd := SetPostInitialSepDigit("user", sep)
		str := strings.Split(spsd, string(sep))
		integer, err := strconv.Atoi(str[0])
		require.NoError(t, err)
		require.Equal(t,
			strconv.Itoa(integer)+string(sep)+"user", spsd)
	})

	t.Run("SepWithRandomDigit", func(t *testing.T) {
		var sep byte = '.'
		spsd := SepWithRandomDigit("user", sep, 10)
		str := strings.Split(spsd, string(sep))
		integer, err := strconv.Atoi(str[len(str)-1])
		require.NoError(t, err)
		require.Equal(t,
			"user"+string(sep)+strconv.Itoa(integer), spsd)

		require.Less(t, integer, 10)

		spsd2n := SepWithRandomDigit("user", sep, 100)
		str2n := strings.Split(spsd2n, string(sep))
		integer2n, err := strconv.Atoi(str2n[len(str2n)-1])
		require.NoError(t, err)

		require.Less(t, integer2n, 100)

		spsd3n := SepWithRandomDigit("user", sep, 1000)
		str3n := strings.Split(spsd3n, string(sep))
		integer3n, err := strconv.Atoi(str3n[len(str3n)-1])
		require.NoError(t, err)

		require.Less(t, integer3n, 1000)
	})

	t.Run("SwapTwoChars", func(t *testing.T) {
		stc := SwapTwoChars("user")
		require.Equal(t, "usre", stc)
		// If the string length less 3 chars
		stcLess3char := SwapTwoChars("ab")
		require.Equal(t, "ab", stcLess3char)
		// If the string length is even
		stcEvenLen := SwapTwoChars("unamee")
		require.Equal(t, "unmaee", stcEvenLen)
	})

	t.Run("RepeatInitialAppendDigit", func(t *testing.T) {
		riad := RepeatInitialAppendDigit("user", 10)
		str := riad[len(riad)-1:]
		integer, err := strconv.Atoi(str)
		require.NoError(t, err)
		require.Equal(t, "uuser"+strconv.Itoa(integer), riad)
	})

	t.Run("PrefixRandomDigit", func(t *testing.T) {
		for n := 100; n < 1001; n += 900 {
			builtN := ""
			prd := PrefixRandomDigit("user", n)
			for i := 0; i < 3; i++ {
				_, err := strconv.Atoi(string(prd[i]))
				if err == nil {
					builtN += string(prd[i])
				}
			}
			require.Equal(t, builtN+"user", prd)
		}
	})

	t.Run("SuffixRandomDigit", func(t *testing.T) {
		srd := SuffixRandomDigit("user", 10)
		str := srd[len(srd)-1:]
		integer, err := strconv.Atoi(str)
		require.NoError(t, err)
		require.Equal(t, "user"+strconv.Itoa(integer), srd)
	})

	t.Run("SetPrefixRandomDigit", func(t *testing.T) {
		// Since SetPrefixRandomDigit runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			sprd := SetPrefixRandomDigit("user")
			builtN := ""
			for i := 0; i < 3; i++ {
				_, err := strconv.Atoi(string(sprd[i]))
				if err == nil {
					builtN += string(sprd[i])
				}
			}
			require.Equal(t, builtN+"user", sprd)
		}

	})

	t.Run("SetSuffixRandomDigit", func(t *testing.T) {
		// Since SetSuffixRandomDigit runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			ssrd := SetSuffixRandomDigit("user")
			builtN := ""
			for _, num := range ssrd[len(ssrd)-3:] {
				_, err := strconv.Atoi(string(num))
				if err == nil {
					builtN += string(num)
				}
			}
			require.Equal(t, "user"+builtN, ssrd)
		}

	})

	t.Run("SetSepWithRandomDigit", func(t *testing.T) {
		// Since SetSuffixRandomDigit runs on randomness,
		// not all cases can be tested directly and
		// we need more iterations to make sure that all
		// possible cases of randomness are covered.
		for n := 0; n < 20; n++ {
			ssrd := SetSepWithRandomDigit("user")
			builtN := ""
			for _, num := range ssrd[len(ssrd)-3:] {
				_, err := strconv.Atoi(string(num))
				if err == nil {
					builtN += string(num)
				}
			}
			require.Equal(t, "user."+builtN, ssrd)
		}
	})

	t.Run("isVowel", func(t *testing.T) {
		falseValue := isVowel('2')
		require.False(t, falseValue)
	})

	t.Run("vowelSwap", func(t *testing.T) {
		vowels := "aeiouAEIOU0"
		for _, v := range vowels {
			out := vowelSwap(byte(v))
			require.NotEqual(t, out, v)
		}
	})

	t.Run("alphabetSwap", func(t *testing.T) {
		vowels := "aeiouAEIOU0ckgjbpfv"
		for _, v := range vowels {
			out := alphabetSwap(byte(v))
			require.NotEqual(t, out, v)
		}
	})

}
