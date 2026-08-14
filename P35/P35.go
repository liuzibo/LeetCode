package p35

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 防止溢出，等同于 (left + right) / 2
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 循环结束后，left 就是插入位置
	return left
}
