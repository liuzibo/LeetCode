package P207

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 建图：邻接表 + 入度数组
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		// pre = [ai, bi]，表示 bi -> ai
		from, to := pre[1], pre[0]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// 2. 将所有入度为 0 的节点入队
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. BFS 拓扑排序
	count := 0 // 记录已处理的节点数
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++

		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 如果所有节点都被处理，说明无环
	return count == numCourses
}
