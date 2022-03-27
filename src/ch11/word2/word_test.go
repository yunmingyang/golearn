package word

import (
	"math/rand"
	"testing"
	"time"
)


func randomPalindrom(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)

	for i := 0; i < (n + 1) / 2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n - 1 - i] = r
	}

	return string(runes)
}

func TestPalindrome(t *testing.T) {
	tests := []struct{
		input string
		want bool
	} {
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrom(test.input); got != test.want {
			t.Errorf(`IsPalindrome("%q") = %v\n`, test.input, !test.want)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %v", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrom(rng)
		if !IsPalindrom(p) {
			t.Errorf(`IsPalindrome("%q") = false`, p)
		}
	}
}

func BenchmarkPalindromes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrom("A man, a plan, a cancl: Panama")
	}
}