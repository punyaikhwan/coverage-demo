package stringutils

// Concat concatenates two strings and returns the result.
func Concat(a, b string) string {
	return a + b
}

// Reverse reverses a string and returns the result.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome.
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

// IsAnagram checks if two strings are anagrams.
// Two strings are anagrams if they contain the same characters in any order.
func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return IsAnagramHelper(a, b)
}

// IsAnagramHelper checks if two strings are anagrams.
func IsAnagramHelper(a, b string) bool {
	counts := make(map[rune]int)
	for _, r := range a {
		counts[r]++
	}
	for _, r := range b {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}
	return true
}
