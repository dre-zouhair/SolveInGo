package cmd

func LongestPalindrome(s string) string {
	longestPalindromeRes := ""

	for i := range s {
		longestPalindromeRes = getPalindrome(s, i, i, longestPalindromeRes)
		longestPalindromeRes = getPalindrome(s, i, i+1, longestPalindromeRes)
	}

	return longestPalindromeRes
}

func getPalindrome(s string, l int, r int, longestPalindromeRes string) string {
	for ; l > -1 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
		if r-l+1 > len(longestPalindromeRes) {
			longestPalindromeRes = s[l:(r + 1)]
		}
	}
	return longestPalindromeRes
}
