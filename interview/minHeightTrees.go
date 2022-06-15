package interview

/**
If the number of nodes is V, and the number of edges is E.
The space complexity is O(V+2E), for storing the whole tree.

The time complexity is O(E),
because we gradually remove all the neighboring information.
As some friends pointing out, for a tree, if V=n, then E=n-1.
Thus both time complexity and space complexity become O(n).
*/
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
