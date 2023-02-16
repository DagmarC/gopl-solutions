package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
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
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func nonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2 // random length up to 2 to 24
	b := make([]byte, n)

	for i := 0; i < (n+1)/2; i++ {
		r := letterBytes[rand.Intn(len(letterBytes))]
		b[i] = r

		var q byte
		// Loop while q == r to prevent having palindrome
		for q = letterBytes[rand.Intn(len(letterBytes))]; strings.EqualFold(string(q), string(r)); {
			q = letterBytes[rand.Intn(len(letterBytes))]
		}
		b[n-1-i] = q // q is not equak to r at this time
	}
	return string(b)
}

//!+random
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		var p string
		if i%2 == 0 {
			p = nonPalindrome(rng)
			if IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = true, expected false", p)
			}
		} else {
			p = randomPalindrome(rng)
			if !IsPalindrome(p) {
				t.Errorf("IsPalindrome(%q) = false, expected true", p)
			}
		}
	}
}
