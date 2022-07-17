package interview

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}
	dfsFillColor(image, sr, sc, image[sr][sc], color)
	return image
}

func dfsFillColor(image [][]int, i, j, oldColor, color int) {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) || image[i][j] != oldColor {
		return
	}

	image[i][j] = color

	dfsFillColor(image, i+1, j, oldColor, color)
	dfsFillColor(image, i-1, j, oldColor, color)
	dfsFillColor(image, i, j+1, oldColor, color)
	dfsFillColor(image, i, j-1, oldColor, color)
}
