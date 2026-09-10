package P300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, 0)

	for _, num := range nums {
		// 二分查找第一个 >= num 的位置
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 情况一：没找到，追加
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// 情况二：找到，替换
			tails[left] = num
		}
	}

	return len(tails)
}
