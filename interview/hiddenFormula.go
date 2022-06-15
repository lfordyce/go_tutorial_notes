package interview

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
func findSolution(customFunction func(int, int) int, z int) [][]int {
	queue := make([][]int, 0)
	x, y := 1, z
	for x <= z && y > 0 {
		cf := customFunction(x, y)
		if cf == z {
			queue = append(queue, []int{x, y})
			x, y = x+1, y-1
		} else if cf > z {
			y -= 1
		} else {
			x += 1
		}
	}
	return queue
}
