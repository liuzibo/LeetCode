package p46

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// isPermutation 判断 p 是否是 nums 的一个合法全排列（包含相同元素，顺序不同）
func isPermutation(nums, p []int) bool {
	if len(nums) != len(p) {
		return false
	}
	a := make([]int, len(nums))
	b := make([]int, len(p))
	copy(a, nums)
	copy(b, p)
	sort.Ints(a)
	sort.Ints(b)
	return reflect.DeepEqual(a, b)
}

// validatePermutations 校验结果集：数量正确、每个排列合法、无重复
func validatePermutations(t *testing.T, nums []int, got [][]int) {
	t.Helper()

	wantCount := factorial(len(nums))
	if len(got) != wantCount {
		t.Errorf("排列数量 = %d, want %d (n!)", len(got), wantCount)
		return
	}

	seen := make(map[string]bool, len(got))
	for i, p := range got {
		if !isPermutation(nums, p) {
			t.Errorf("第 %d 个排列 %v 不是 %v 的合法全排列", i, p, nums)
			return
		}
		key := keyOf(p)
		if seen[key] {
			t.Errorf("存在重复排列: %v", p)
			return
		}
		seen[key] = true
	}
}

func factorial(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func keyOf(p []int) string {
	s := ""
	for i, v := range p {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%d", v)
	}
	return s
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "LeetCode官方示例",
			nums: []int{1, 2, 3},
		},
		{
			name: "单元素",
			nums: []int{0},
		},
		{
			name: "两元素",
			nums: []int{1, 2},
		},
		{
			name: "空数组",
			nums: []int{},
		},
		{
			name: "四元素",
			nums: []int{1, 2, 3, 4},
		},
		{
			name: "含负数",
			nums: []int{-1, 0, 1},
		},
		{
			name: "非连续值",
			nums: []int{5, 10, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			validatePermutations(t, tt.nums, got)
		})
	}
}

func TestPermute_SpecificCases(t *testing.T) {
	t.Run("单元素结果", func(t *testing.T) {
		got := permute([]int{5})
		want := [][]int{{5}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("permute([5]) = %v, want %v", got, want)
		}
	})

	t.Run("空数组结果", func(t *testing.T) {
		got := permute([]int{})
		want := [][]int{{}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("permute([]) = %v, want %v", got, want)
		}
	})

	t.Run("两元素结果", func(t *testing.T) {
		got := permute([]int{1, 2})
		want := [][]int{{1, 2}, {2, 1}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("permute([1,2]) = %v, want %v", got, want)
		}
	})
}
