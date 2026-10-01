package P1

import "testing"

// TestTwoSum 以表驱动方式测试 twoSum。
//
// 说明:实现返回的下标顺序为「后遍历到的下标在前」,例如示例1返回 [1 0]。
// LeetCode 官方判定下标顺序不限,因此这里对期望结果做无序比较。
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int // 期望的下标对;无解时为 nil
	}{
		{
			name:   "示例1:基础用例",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "示例2:答案不含下标0",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "示例3:两个重复元素",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "负数",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "包含0,答案由两个0组成",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "答案位于数组两端",
			nums:   []int{2, 3, 6},
			target: 8,
			want:   []int{0, 2},
		},
		{
			name:   "无解返回nil",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   nil,
		},
		{
			name:   "空数组返回nil",
			nums:   []int{},
			target: 5,
			want:   nil,
		},
		{
			name:   "不能重复使用同一元素",
			nums:   []int{3, 5},
			target: 6,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			// 无解场景:期望返回 nil
			if tt.want == nil {
				if got != nil {
					t.Errorf("twoSum(%v, %d) = %v, 期望返回 nil", tt.nums, tt.target, got)
				}
				return
			}

			if len(got) != 2 {
				t.Fatalf("twoSum(%v, %d) = %v, 期望恰好返回两个下标 %v", tt.nums, tt.target, got, tt.want)
			}

			// 下标顺序不限:正序或反序均视为匹配
			match := (got[0] == tt.want[0] && got[1] == tt.want[1]) ||
				(got[0] == tt.want[1] && got[1] == tt.want[0])
			if !match {
				t.Errorf("twoSum(%v, %d) = %v, 期望下标 %v(顺序不限)", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
