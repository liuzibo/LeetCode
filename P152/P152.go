package P152

func maxProduct(nums []int) int {
	maxProd, minProd, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(cur, maxProd*cur)
		minProd = min(cur, minProd*cur)
		result = max(result, maxProd)
	}
	return result
}
