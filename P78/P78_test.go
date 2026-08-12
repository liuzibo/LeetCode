package p78

import (
	"fmt"
	"reflect"
	"testing"
)

// isSubsequence 判断 sub 是否是 nums 的子序列（元素相对顺序与 nums 一致）
func isSubsequence(nums, sub []int) bool {
	i := 0
	for _, v := range sub {
		for i < len(nums) && nums[i] != v {
			i++
		}
		if i == len(nums) {
			return false
		}
		i++ // 消费掉 nums[i]
	}
	return true
}

// validateSubsets 校验结果集：数量正确（2^n）、每个子集合法、无重复
func validateSubsets(t *testing.T, nums []int, got [][]int) {
	t.Helper()

	wantCount := 1 << len(nums) // 2^n
	if len(got) != wantCount {
		t.Errorf("子集数量 = %d, want %d (2^n)", len(got), wantCount)
		return
	}

	seen := make(map[string]bool, len(got))
	for i, s := range got {
		if !isSubsequence(nums, s) {
			t.Errorf("第 %d 个子集 %v 不是 %v 的合法子序列", i, s, nums)
			return
		}
		key := keyOf(s)
		if seen[key] {
			t.Errorf("存在重复子集: %v", s)
			return
		}
		seen[key] = true
	}
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

func TestSubsets(t *testing.T) {
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
			got := subsets(tt.nums)
			validateSubsets(t, tt.nums, got)
		})
	}
}

func TestSubsets_SpecificCases(t *testing.T) {
	t.Run("空数组结果", func(t *testing.T) {
		got := subsets([]int{})
		want := [][]int{{}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("subsets([]) = %v, want %v", got, want)
		}
	})

	t.Run("单元素结果", func(t *testing.T) {
		got := subsets([]int{5})
		want := [][]int{{}, {5}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("subsets([5]) = %v, want %v", got, want)
		}
	})

	t.Run("两元素结果", func(t *testing.T) {
		got := subsets([]int{1, 2})
		want := [][]int{{}, {1}, {1, 2}, {2}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("subsets([1,2]) = %v, want %v", got, want)
		}
	})

	t.Run("三元素结果", func(t *testing.T) {
		got := subsets([]int{1, 2, 3})
		want := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("subsets([1,2,3]) = %v, want %v", got, want)
		}
	})
}
