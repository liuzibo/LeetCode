package P209

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	minLen := len(nums) + 1

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minLen = min(right-left+1, minLen)
			sum -= nums[left]
			left++
		}
		right++
	}
	if minLen == len(nums)+1 {
		minLen = 0
	}
	return minLen
}
