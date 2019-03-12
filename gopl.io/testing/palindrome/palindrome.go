// Package  palindrome provides utilities for word games
package palindrome

// Is Palindrome reports whether s reads the same forward and backward.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
