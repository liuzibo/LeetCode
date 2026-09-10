package P4

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	totalLeft := (m + n + 1) / 2
	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := totalLeft - i

		nums1LeftMax := math.MinInt32
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}
		nums1RightMin := math.MaxInt32
		if i < m {
			nums1RightMin = nums1[i]
		}
		nums2LeftMax := math.MinInt32
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}
		nums2RightMin := math.MaxInt32
		if j < n {
			nums2RightMin = nums2[j]
		}

		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			if (m+n)%2 == 1 {
				return float64(max(nums1LeftMax, nums2LeftMax))
			} else {
				return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
			}
		} else if nums1LeftMax > nums2RightMin {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return 0.0
}
