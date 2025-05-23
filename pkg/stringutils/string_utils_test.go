package stringutils

import "testing"

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "empty strings",
			a:        "",
			b:        "",
			expected: "",
		},
		{
			name:     "first string empty",
			a:        "",
			b:        "world",
			expected: "world",
		},
		{
			name:     "second string empty",
			a:        "hello",
			b:        "",
			expected: "hello",
		},
		{
			name:     "both strings non-empty",
			a:        "hello ",
			b:        "world",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Concat(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Concat(%q, %q) = %q, expected %q", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "palindrome",
			input:    "radar",
			expected: "radar",
		},
		{
			name:     "non-palindrome",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "with unicode",
			input:    "こんにちは",
			expected: "はちにんこ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "palindrome - even length",
			input:    "abba",
			expected: true,
		},
		{
			name:     "palindrome - odd length",
			input:    "radar",
			expected: true,
		},
		{
			name:     "not a palindrome",
			input:    "hello",
			expected: false,
		},
		{
			name:     "unicode palindrome",
			input:    "あいういあ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
