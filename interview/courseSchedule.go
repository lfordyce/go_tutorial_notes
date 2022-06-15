package interview

/**
here are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return true if you can finish all courses. Otherwise, return false.
*/

// build adjacency list from edges which will allow us
// to perform common operations like getting all descendants
// of a given vertex in 0(1)
func canFinish(n int, prerequisites [][]int) bool {
	// a list to store the number of incoming edges of each vertex
	in := make([]int, n)
	frees := make([][]int, n)
	queue := make([]int, 0, n)
	// build adjacency list
	for _, v := range prerequisites {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}
	// a queue of all vertices with no incoming edge
	// at least one such node must exist in a non-empty acyclic graph
	// vertices in this queue have the same order as the eventual topological sort
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	for i := 0; i != len(queue); i++ {
		// pop vertex from front of queue
		c := queue[i]
		v := frees[c]
		// for each descendant of current vertx reduce its in-degree by 1
		for _, vv := range v {
			in[vv]--
			// if in-degree becomes 0, add it to the queue
			if in[vv] == 0 {
				queue = append(queue, vv)
			}
		}
	}
	return len(queue) == n
}
