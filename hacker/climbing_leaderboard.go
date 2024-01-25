package hacker

import "sort"

func climbinLeaderboard(scores []int32, alice []int32) []int32 {
	var pos int32 = 1
	currentIndex := 0
	results := make([]int32, len(alice))

	for i := len(alice) - 1; i >= 0; i-- {
		aliceScore := alice[i]
		for j := currentIndex; j < len(scores); j++ {
			score := scores[j]
			if aliceScore >= score {
				results[i] = pos
				currentIndex = j
				break
			} else if j == len(scores)-1 {
				for k := i; k <= 0; k++ {
					results[k] = pos + 1
				}
				return results
			} else if score > scores[j+1] {
				pos++
			}
			currentIndex = j
		}
	}

	return results
}

func climbingLeaderboardAlt(scores []int32, alice []int32) []int32 {
	ranks := scores[:1]
	last := ranks[0]
	for _, score := range scores[1:] {
		if last != score {
			ranks = append(ranks, score)
		}
		last = score
	}
	results := make([]int32, 0, len(alice))
	for _, score := range alice {
		rank := sort.Search(len(ranks), func(i int) bool { return ranks[i] <= score })
		results = append(results, int32(rank+1))
	}
	return results
}
