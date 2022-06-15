package interview

/**
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree). Last year, The ministry of transport decided to orient the roads in one direction because they are too narrow.

Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.

This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.

It's guaranteed that each city can reach city 0 after reorder.
*/
func minReorder(n int, connections [][]int) int {
	adjList := make([][]int, n)
	for _, c := range connections {
		adjList[c[0]] = append(adjList[c[0]], c[1])
		adjList[c[1]] = append(adjList[c[1]], -c[0])
	}
	return dfs(adjList, make([]bool, n), 0)
}

func dfs(adjList [][]int, visited []bool, from int) int {
	change := 0
	visited[from] = true
	for _, to := range adjList[from] {
		if !visited[abs(to)] {
			change += dfs(adjList, visited, abs(to))
			if to > 0 {
				change += 1
			}
		}
	}
	return change
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
