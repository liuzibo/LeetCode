package p39

import (
	"fmt"
	"sort"
	"testing"
)

// allFromCandidates 判断 comb 中的每个元素都来自 candidates
func allFromCandidates(candidates, comb []int) bool {
	set := make(map[int]bool, len(candidates))
	for _, c := range candidates {
		set[c] = true
	}
	for _, v := range comb {
		if !set[v] {
			return false
		}
	}
	return true
}

// sumOf 返回切片元素之和
func sumOf(comb []int) int {
	s := 0
	for _, v := range comb {
		s += v
	}
	return s
}

// normalizedKey 返回排序后的组合字符串键，用于去重比较
func normalizedKey(comb []int) string {
	a := make([]int, len(comb))
	copy(a, comb)
	sort.Ints(a)
	s := ""
	for i, v := range a {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf("%d", v)
	}
	return s
}

// validateCombinations 校验结果集：每个组合之和等于 target、元素均来自 candidates、无重复组合
func validateCombinations(t *testing.T, candidates []int, target int, got [][]int) {
	t.Helper()

	seen := make(map[string]bool, len(got))
	for i, comb := range got {
		if sumOf(comb) != target {
			t.Errorf("第 %d 个组合 %v 之和 = %d, want %d", i, comb, sumOf(comb), target)
			return
		}
		if !allFromCandidates(candidates, comb) {
			t.Errorf("第 %d 个组合 %v 含有不在 candidates %v 中的元素", i, comb, candidates)
			return
		}
		key := normalizedKey(comb)
		if seen[key] {
			t.Errorf("存在重复组合: %v", comb)
			return
		}
		seen[key] = true
	}
}

// expectCombinations 比对实际结果与期望结果（顺序无关，按归一化键匹配）
func expectCombinations(t *testing.T, got, want [][]int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("组合数量 = %d, want %d\ngot: %v\nwant: %v", len(got), len(want), got, want)
		return
	}
	wantSet := make(map[string]bool, len(want))
	for _, c := range want {
		wantSet[normalizedKey(c)] = true
	}
	for _, c := range got {
		key := normalizedKey(c)
		if !wantSet[key] {
			t.Errorf("未预期的组合 %v", c)
			return
		}
	}
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "LeetCode官方示例1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "LeetCode官方示例2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "LeetCode官方示例3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "单候选数恰好命中",
			candidates: []int{1},
			target:     1,
			want:       [][]int{{1}},
		},
		{
			name:       "单候选数多次重复",
			candidates: []int{1},
			target:     3,
			want:       [][]int{{1, 1, 1}},
		},
		{
			name:       "目标为0（空组合）",
			candidates: []int{2, 3},
			target:     0,
			want:       [][]int{{}},
		},
		{
			name:       "无解（候选数均大于目标",
			candidates: []int{5, 6, 7},
			target:     3,
			want:       [][]int{},
		},
		{
			name:       "多种组合去重校验",
			candidates: []int{2, 3, 5},
			target:     7,
			want:       [][]int{{2, 2, 3}, {2, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			validateCombinations(t, tt.candidates, tt.target, got)
			expectCombinations(t, got, tt.want)
		})
	}
}

func TestCombinationSum_SpecificCases(t *testing.T) {
	t.Run("示例1结果精确比对", func(t *testing.T) {
		got := combinationSum([]int{2, 3, 6, 7}, 7)
		want := [][]int{{2, 2, 3}, {7}}
		if len(got) != len(want) {
			t.Errorf("组合数量 = %d, want %d", len(got), len(want))
			return
		}
		// 逐个比对（顺序无关）
		gotSet := make(map[string]bool)
		for _, c := range got {
			gotSet[normalizedKey(c)] = true
		}
		for _, c := range want {
			if !gotSet[normalizedKey(c)] {
				t.Errorf("缺少期望组合 %v, got %v", c, got)
			}
		}
	})

	t.Run("无解返回空", func(t *testing.T) {
		got := combinationSum([]int{2}, 1)
		if len(got) != 0 {
			t.Errorf("combinationSum([2], 1) = %v, want []", got)
		}
	})

	t.Run("目标0返回单个空组合", func(t *testing.T) {
		got := combinationSum([]int{2, 3}, 0)
		// 注意：函数返回的内层切片可能为 nil（append(nil) 无元素时为 nil），
		// 与 [][]int{{}} 在 reflect.DeepEqual 下不等，因此只校验数量与元素长度。
		if len(got) != 1 {
			t.Fatalf("组合数量 = %d, want 1, got %v", len(got), got)
		}
		if len(got[0]) != 0 {
			t.Errorf("组合元素 = %v, want 空组合", got[0])
		}
	})

	t.Run("含较大候选数", func(t *testing.T) {
		candidates := []int{7, 14, 21}
		target := 28
		got := combinationSum(candidates, target)
		validateCombinations(t, candidates, target, got)
		// 28 = 14+14 = 7+21 = 7+7+14 = 7+7+7+7
		if len(got) != 4 {
			t.Errorf("组合数量 = %d, want 4, got %v", len(got), got)
		}
	})
}
