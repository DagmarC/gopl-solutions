package main

import (
	"math/rand"
	"testing"
	"time"
)

//!+random
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to '\u099'
		runes[i] = r
		runes[n-1-i] = r
	}
	return "? " + string(runes) + "@" // handling of Punctuation and spaces
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
