package p79

import (
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "LeetCode官方示例1-存在",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "LeetCode官方示例2-存在",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "SEE",
			want:  true,
		},
		{
			name:  "LeetCode官方示例3-不存在",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
		{
			name:  "单字符匹配",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
		{
			name:  "单字符不匹配",
			board: [][]byte{{'A'}},
			word:  "B",
			want:  false,
		},
		{
			name:  "空board",
			board: [][]byte{},
			word:  "A",
			want:  false,
		},
		{
			name:  "空行board",
			board: [][]byte{{}},
			word:  "A",
			want:  false,
		},
		{
			name:  "单词长度超过board单元格数",
			board: [][]byte{{'A', 'B'}, {'C', 'D'}},
			word:  "ABCDEF",
			want:  false,
		},
		{
			name:  "长路径蛇形遍历",
			board: [][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}},
			word:  "ABCFEDGHI",
			want:  true,
		},
		{
			name:  "需要回溯-首选方向走入死路",
			board: [][]byte{{'A', 'B', 'C', 'D'}, {'B', 'X', 'X', 'E'}},
			word:  "ABCD",
			want:  true,
		},
		{
			name:  "不能重复使用同一单元格",
			board: [][]byte{{'A', 'B'}, {'B', 'A'}},
			word:  "ABA",
			want:  true,
		},
		{
			name:  "不能重复使用同一单元格-失败",
			board: [][]byte{{'A', 'A'}},
			word:  "AAA",
			want:  false,
		},
		{
			name:  "水平遍历",
			board: [][]byte{{'A', 'B', 'C', 'D'}},
			word:  "DCBA",
			want:  true,
		},
		{
			name:  "垂直遍历",
			board: [][]byte{{'A'}, {'B'}, {'C'}, {'D'}},
			word:  "ABCD",
			want:  true,
		},
		{
			name:  "首字符不存在",
			board: [][]byte{{'A', 'B'}, {'C', 'D'}},
			word:  "XYZ",
			want:  false,
		},
		{
			name:  "长单词成功",
			board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			word:  "oath",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
