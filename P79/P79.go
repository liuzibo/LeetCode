package p79

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 四个方向：上、下、左、右
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		// 当前字符不匹配
		if board[x][y] != word[idx] {
			return false
		}
		// 已匹配到最后一个字符
		if idx == len(word)-1 {
			return true
		}

		visited[x][y] = true
		defer func() { visited[x][y] = false }() // 回溯

		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				if dfs(nx, ny, idx+1) {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
