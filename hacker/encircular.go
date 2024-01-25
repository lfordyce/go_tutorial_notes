package hacker

func doesCircleExist(commands []string) []string {
	// Write your code here

	var output []string
	for _, cmd := range commands {
		if isRobotBounded(cmd) {
			output = append(output, "YES")
		} else {
			output = append(output, "NO")
		}
	}
	return output
}

func isRobotBounded(instructions string) bool {
	// robot is bounded if:
	// end position is 0,0
	// OR
	// end direction is not North
	direction := 'N'
	x := 0
	y := 0
	for _, c := range instructions {
		if c == 'G' {
			if direction == 'N' {
				y++
			} else if direction == 'S' {
				y--
			} else if direction == 'W' {
				x--
			} else {
				x++
			}
		} else if c == 'L' {
			if direction == 'N' {
				direction = 'W'
			} else if direction == 'S' {
				direction = 'E'
			} else if direction == 'W' {
				direction = 'S'
			} else {
				direction = 'N'
			}
		} else {
			if direction == 'N' {
				direction = 'E'
			} else if direction == 'S' {
				direction = 'W'
			} else if direction == 'W' {
				direction = 'N'
			} else {
				direction = 'S'
			}
		}
	}
	return (x == 0 && y == 0) || (direction != 'N')
}
