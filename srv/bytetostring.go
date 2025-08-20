package srv

func IsNumberPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	return n == reverse(n)
}

func reverse(n int) int {
	var reversed int
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}
