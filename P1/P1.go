package P1

func twoSum(nums []int, target int) []int {
	length := len(nums)
	// key是nums[i], value是i
	hashMap := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if j, ok := hashMap[target-nums[i]]; ok {
			return []int{i, j}
		}
		hashMap[nums[i]] = i
	}
	return nil
}
