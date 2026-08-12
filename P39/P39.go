package p39

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	// 排序便于剪枝
	sort.Ints(candidates)

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
		// 找到一组有效组合
		if target == 0 {
			// 需要复制 path，因为后续回溯会修改 path
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// 剪枝：如果 target < 0 直接返回
		if target < 0 {
			return
		}

		for i := start; i < len(candidates) && candidates[i] <= target; i++ {
			// 选择当前数字
			path = append(path, candidates[i])
			// 关键：下一层递归仍然从 i 开始，因为可以重复选取同一个数字
			backtrack(i, target-candidates[i])
			// 回溯，撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return res
}
