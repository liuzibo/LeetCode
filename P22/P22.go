package p22

func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(path string, left, right int)
	backtrack = func(path string, left, right int) {
		// 终止条件：左右括号都用完了
		if len(path) == 2*n {
			res = append(res, path)
			return
		}

		// 分支1：还可以加左括号
		if left < n {
			backtrack(path+"(", left+1, right)
		}

		// 分支2：右括号数量少于左括号时才能加右括号
		if right < left {
			backtrack(path+")", left, right+1)
		}
	}

	backtrack("", 0, 0)
	return res
}
