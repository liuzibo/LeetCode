package p78

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 将当前路径的副本加入结果集
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		// 从 start 开始，依次尝试添加每个元素
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1] // 回溯，撤销选择
		}
	}

	backtrack(0)
	return res
}
