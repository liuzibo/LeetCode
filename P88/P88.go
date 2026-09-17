package P88

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)
	i, j := 0, 0
	cur := 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			res[cur] = nums1[i]
			cur++
			i++
		} else {
			res[cur] = nums2[j]
			cur++
			j++
		}
	}
	for i < m {
		res[cur] = nums1[i]
		cur++
		i++
	}
	for j < n {
		res[cur] = nums2[j]
		cur++
		j++
	}
	copy(nums1, res)
}
