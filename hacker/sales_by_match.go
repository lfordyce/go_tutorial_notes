package hacker

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	pairMap := make(map[int32]struct{}, 0)

	var pairs int32
	for i := 0; i < len(ar); i++ {
		if _, ok := pairMap[ar[i]]; !ok {
			pairMap[ar[i]] = struct{}{}
		} else {
			pairs++
			delete(pairMap, ar[i])
		}
	}
	return pairs
}

func sockMerchantAlt(n int32, ar []int32) int32 {
	pairsToSell := make([]int32, 0, 9)
	var countToSell int32
	for _, sock := range ar {
		ind := sock % n
		if pairsToSell[ind] == 2 {
			countToSell++
			pairsToSell[ind] = 0
		}
	}
	return countToSell
}