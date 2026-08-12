package p46

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		// 递归终止条件：path 长度等于 nums 长度，说明找到了一个全排列
		if len(path) == len(nums) {
			// 注意：需要 copy 一份 path 的副本，因为后续回溯会修改 path
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 做选择
			used[i] = true
			path = append(path, nums[i])
			// 递归进入下一层
			backtrack()
			// 撤销选择（回溯）
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
