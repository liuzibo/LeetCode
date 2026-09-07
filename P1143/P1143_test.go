package P1143

import "testing"

// TestLongestCommonSubsequence 使用"表驱动测试"（table-driven test），
// 这是 Go 中最常用的测试写法：把所有测试用例放在一个切片里，循环依次执行。
func TestLongestCommonSubsequence(t *testing.T) {
	// tests 是一个匿名结构体切片，每个元素代表一个测试用例
	// name: 用例的名字（失败时方便定位是哪个用例出错了）
	// text1, text2: 输入参数
	// want: 期望的返回值
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "LeetCode官方示例1",
			text1: "abcde",
			text2: "ace",
			want:  3, // 最长公共子序列是 "ace"
		},
		{
			name:  "LeetCode官方示例2",
			text1: "abc",
			text2: "abc",
			want:  3, // 两个字符串完全相同
		},
		{
			name:  "LeetCode官方示例3",
			text1: "abc",
			text2: "def",
			want:  0, // 没有任何公共字符
		},
		{
			name:  "空字符串",
			text1: "",
			text2: "abc",
			want:  0, // 边界情况：空串和任何串的公共子序列长度为 0
		},
		{
			name:  "部分匹配",
			text1: "bsbininm",
			text2: "jmjkbkjkv",
			want:  1, // 只有 'b' 一个公共字符
		},
	}

	// 遍历每个用例，t.Run 会把每个用例作为子测试独立运行
	// 好处：单个用例失败不影响其他用例，且能用 go test -run 命令单独跑某个用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 调用被测函数，拿到实际结果
			got := longestCommonSubsequence(tt.text1, tt.text2)

			// 比较实际值和期望值，不相等就报告错误
			// t.Errorf 不会中断测试，会继续执行后面的用例（t.Fatalf 才会立即中断）
			if got != tt.want {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d",
					tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}
