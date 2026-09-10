package P763

func partitionLabels(s string) []int {
	last := make([]int, 26)
	for i, ch := range s {
		last[ch-'a'] = i
	}

	var ans []int
	start, right := 0, 0
	for i, ch := range s {
		right = max(right, last[ch-'a'])
		if i == right {
			ans = append(ans, right-start+1)
			start = right + 1
		}
	}
	return ans
}
