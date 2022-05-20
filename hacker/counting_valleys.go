package hacker

func countingValleys(steps int32, path string) int32 {

	var vallyCounter, altitude int32

	for _, char := range path {
		//ch := string(char)
		if char == 'U' {
			altitude++
			if altitude == 0 {
				vallyCounter++
			}
		} else {
			altitude--
		}
	}
	return vallyCounter
}