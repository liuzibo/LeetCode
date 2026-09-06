package P5

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	start, maxLen := 0, 1
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] && length == 2 {
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
			if dp[i][j] && length > maxLen {
				start = i
				maxLen = length
			}
		}
	}
	return s[start : start+maxLen]
}
