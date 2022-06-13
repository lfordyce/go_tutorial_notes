package hacker

import "sort"

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, budget int32) int32 {
	// Write your code here
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	count := int32(0)
	for _, price := range prices {
		if price <= budget {
			budget -= price
			count++
		}
	}

	return count
}
