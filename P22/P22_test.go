package p22

import (
	"sort"
	"testing"
)

// isValidParentheses 校验 s 是否是合法的括号串：长度为 2n 且任意前缀中 '(' 数量 >= ')' 数量
func isValidParentheses(s string, n int) bool {
	if len(s) != 2*n {
		return false
	}
	bal := 0
	for _, c := range s {
		if c == '(' {
			bal++
		} else if c == ')' {
			bal--
		} else {
			return false
		}
		if bal < 0 {
			return false
		}
	}
	return bal == 0
}

// catalan 返回第 n 个卡特兰数 C(n) = C(2n,n)/(n+1)
func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}

// validateParentheses 校验结果集：数量正确(C(n))、每个串合法、无重复
func validateParentheses(t *testing.T, n int, got []string) {
	t.Helper()

	wantCount := catalan(n)
	if len(got) != wantCount {
		t.Errorf("括号组合数量 = %d, want %d (Catalan(%d))", len(got), wantCount, n)
		return
	}

	seen := make(map[string]bool, len(got))
	for i, s := range got {
		if !isValidParentheses(s, n) {
			t.Errorf("第 %d 个组合 %q 不是合法的括号串", i, s)
			return
		}
		if seen[s] {
			t.Errorf("存在重复组合: %q", s)
			return
		}
		seen[s] = true
	}
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "LeetCode官方示例",
			n:    3,
		},
		{
			name: "n=1",
			n:    1,
		},
		{
			name: "n=2",
			n:    2,
		},
		{
			name: "n=4",
			n:    4,
		},
		{
			name: "n=5",
			n:    5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			validateParentheses(t, tt.n, got)
		})
	}
}

func TestGenerateParenthesis_SpecificCases(t *testing.T) {
	t.Run("n=1结果", func(t *testing.T) {
		got := generateParenthesis(1)
		want := []string{"()"}
		if !sortedEqual(got, want) {
			t.Errorf("generateParenthesis(1) = %v, want %v", got, want)
		}
	})

	t.Run("n=2结果", func(t *testing.T) {
		got := generateParenthesis(2)
		want := []string{"(())", "()()"}
		if !sortedEqual(got, want) {
			t.Errorf("generateParenthesis(2) = %v, want %v", got, want)
		}
	})

	t.Run("n=3结果", func(t *testing.T) {
		got := generateParenthesis(3)
		want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		if !sortedEqual(got, want) {
			t.Errorf("generateParenthesis(3) = %v, want %v", got, want)
		}
	})

	t.Run("n=4数量校验", func(t *testing.T) {
		got := generateParenthesis(4)
		if len(got) != 14 {
			t.Errorf("generateParenthesis(4) 数量 = %d, want 14", len(got))
		}
	})
}

// sortedEqual 比较两个字符串切片，排序后是否相等
func sortedEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	ca := make([]string, len(a))
	cb := make([]string, len(b))
	copy(ca, a)
	copy(cb, b)
	sort.Strings(ca)
	sort.Strings(cb)
	for i := range ca {
		if ca[i] != cb[i] {
			return false
		}
	}
	return true
}
