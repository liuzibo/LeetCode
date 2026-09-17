package P27

func removeElement(nums []int, val int) int {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}
