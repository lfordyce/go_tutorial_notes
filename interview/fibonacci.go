package interview

import "math/big"

func fibonacciBad(n uint) uint {
	if n < 2 {
		return n
	}

	return fibonacciBad(n-1) + fibonacciBad(n-2)
}

var (
	fibonacciCache = make(map[uint]uint)
)

func fibonacciMemo(n uint) uint {
	if n < 2 {
		return n
	}

	if result, ok := fibonacciCache[n]; ok {
		return result
	}
	fib := fibonacciMemo(n-2) + fibonacciMemo(n-1)
	fibonacciCache[n] = fib
	return fib
}

func fibonacciIter(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint

	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}

func fibonacciBig(n uint) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)

	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}
