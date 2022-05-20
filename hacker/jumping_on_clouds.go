package hacker

// Complete the jumpingOnClouds function below.
// There is a new mobile game that starts with consecutively numbered clouds.
// Some of the clouds are thunderheads and others are cumulus.
// The player can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus 1 or 2.
// The player must avoid the thunderheads.
// Determine the minimum number of jumps it will take to jump from the starting position to the last cloud.
// It is always possible to win the game.
// For each game, you will get an array of clouds numbered
// if they are safe or if they must be avoided.
func jumpingOnClouds(c []int32) int32 {
	var jumps int32
	for i := 0; i < len(c)-1; i++ {
		jumps++
		if i+2 < len(c) && c[i+2] == 0 {
			i++
		}
	}
	return jumps
}
