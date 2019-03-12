// Package  palindrome provides utilities for word games
package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindromede") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}

}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("ete") = false`)
	}
}
