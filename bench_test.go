package unamex

import (
	"testing"
)

// S = Sequential
// P = Parallel
var username = "someuser"

func BenchmarkS_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(username)
	}
}

func BenchmarkP_New(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			New(username)
		}
	})
}

func BenchmarkS_Suggest(b *testing.B) {
	u := New(username)
	for i := 0; i < b.N; i++ {
		u.Suggest(1)
	}
}

func BenchmarkP_Suggest(b *testing.B) {
	u := New(username)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			u.Suggest(1)
		}
	})
}

func BenchmarkS_dSuggestors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defaultSuggestors()
	}
}

func BenchmarkP_dSuggestors(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			defaultSuggestors()
		}
	})
}

func BenchmarkS_DoubleByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoubleByte('c')
	}
}

func BenchmarkP_DoubleByte(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			DoubleByte('c')
		}
	})
}

func BenchmarkS_VanishVowel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VanishVowel(username)
	}
}

func BenchmarkP_VanishVowel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			VanishVowel(username)
		}
	})
}

func BenchmarkS_VowelTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VowelTransform(username, vowelSwap)
	}
}

func BenchmarkP_VowelTransform(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			VowelTransform(username, vowelSwap)
		}
	})
}

func BenchmarkS_RepeatVowel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatVowel(username)
	}
}

func BenchmarkP_RepeatVowel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RepeatVowel(username)
		}
	})
}

func BenchmarkS_AlphabetTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphabetTransform(username, alphabetSwap)
	}
}

func BenchmarkP_AlphabetTransform(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			AlphabetTransform(username, alphabetSwap)
		}
	})
}

func BenchmarkS_RepeatSubfix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatSubfix(username)
	}
}

func BenchmarkP_RepeatSubfix(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RepeatSubfix(username)
		}
	})
}

func BenchmarkS_RepeatPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatPrefix(username)
	}
}

func BenchmarkP_RepeatPrefix(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RepeatPrefix(username)
		}
	})
}

func BenchmarkS_RepeatSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatSuffix(username)
	}
}

func BenchmarkP_RepeatSuffix(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RepeatSuffix(username)
		}
	})
}

func BenchmarkS_SetPostInitialSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPostInitialSep(username, '.')
	}
}

func BenchmarkP_SetPostInitialSep(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetPostInitialSep(username, '.')
		}
	})
}

func BenchmarkS_SetPenultimateSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPenultimateSep(username, '.')
	}
}

func BenchmarkP_SetPenultimateSep(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetPenultimateSep(username, '.')
		}
	})
}

func BenchmarkS_SetPenultimateSepDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPenultimateSepDigit(username, '.')
	}
}

func BenchmarkP_SetPenultimateSepDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetPenultimateSepDigit(username, '.')
		}
	})
}

func BenchmarkS_SetPostInitialSepDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPostInitialSepDigit(username, '.')
	}
}

func BenchmarkP_SetPostInitialSepDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetPostInitialSepDigit(username, '.')
		}
	})
}

func BenchmarkS_SepWithRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SepWithRandomDigit(username, '.', 2)
	}
}

func BenchmarkP_SepWithRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SepWithRandomDigit(username, '.', 2)
		}
	})
}

func BenchmarkS_SwapTwoChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapTwoChars(username)
	}
}

func BenchmarkP_SwapTwoChars(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SwapTwoChars(username)
		}
	})
}

func BenchmarkS_RepeatInitialAppendDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatInitialAppendDigit(username, 2)
	}
}

func BenchmarkP_RepeatInitialAppendDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			RepeatInitialAppendDigit(username, 2)
		}
	})
}

func BenchmarkS_PrefixRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrefixRandomDigit(username, 2)
	}
}

func BenchmarkP_PrefixRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			PrefixRandomDigit(username, 2)
		}
	})
}

func BenchmarkS_SuffixRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SuffixRandomDigit(username, 2)
	}
}

func BenchmarkP_SuffixRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SuffixRandomDigit(username, 2)
		}
	})
}

func BenchmarkS_SetPrefixRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetPrefixRandomDigit(username)
	}
}

func BenchmarkP_SetPrefixRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetPrefixRandomDigit(username)
		}
	})
}

func BenchmarkS_SetSuffixRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetSuffixRandomDigit(username)
	}
}

func BenchmarkP_SetSuffixRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetSuffixRandomDigit(username)
		}
	})
}

func BenchmarkS_SetSepWithRandomDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetSepWithRandomDigit(username)
	}
}

func BenchmarkP_SetSepWithRandomDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			SetSepWithRandomDigit(username)
		}
	})
}

func BenchmarkS_shuffleSuggestors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffleSuggestors(defaultSuggestors())
	}
}

func BenchmarkP_shuffleSuggestors(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			shuffleSuggestors(defaultSuggestors())
		}
	})
}

func BenchmarkS_isLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isLetter('c')
	}
}

func BenchmarkP_isLetter(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			isLetter('c')
		}
	})
}

func BenchmarkS_isDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isDigit('c')
	}
}

func BenchmarkP_isDigit(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			isDigit('c')
		}
	})
}

func BenchmarkS_isVowel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isVowel('c')
	}
}

func BenchmarkP_isVowel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			isVowel('c')
		}
	})
}
