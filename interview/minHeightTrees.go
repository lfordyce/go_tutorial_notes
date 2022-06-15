package interview

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}
	var res []int
	degrees := make([]int, n)
	adjList := make([][]int, n)
	queue := make([]int, 0, n)
	// create adjacent list
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
		// update how many edges each node has
		degrees[edge[1]]++
		degrees[edge[0]]++
	}
	for i := 0; i < n; i++ {
		// adding all the leaf nodes
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		s := len(queue)
		res = res[:0] //clear the root nodes
		for i := 0; i < s; i++ {
			// pop vertex from front of queue
			c := queue[0]
			queue = queue[1:]
			v := adjList[c]
			// for each descendant of current vertx reduce its in-degree by 1
			res = append(res, c)
			for _, neighbor := range v {
				// decrease degree of neighbour nodes and push leaf nodes into queue
				degrees[neighbor]--
				// if in-degree becomes 1, add it to the queue
				if degrees[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return res
}
