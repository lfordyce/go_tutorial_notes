package hacker

import "strings"

// Complete the repeatedString function below.
// example:
// s := "abcac"
// n := 10
// n % s =>


// def repeatedString(s, n):
//    c = s.count('a')
//    div=n//len(s)
//    if n%len(s)==0:
//        c= c*div
//    else:
//        m = n%len(s)
//        c= c*div+s[:m].count('a')
//    return c
func repeatedString(s string, n int64) int64 {
	//factor := n / int64(len(s))
	rem := n % int64(len(s))
	// ca: remaining 'a' characters
	// c: total amount of characters in the given string without the remaining 'rem'
	var ca, c int64
	for i := int64(len(s)) - 1; i > 0; i-- {
		if s[i] == 'a' {
			c++
			if i < rem {
				ca++
			}
		}
	}
	return ((n - rem) / int64(len(s)) * c) + ca
}

func repeatedStringAlt(s string, n int64) int64 {
	factor := n / int64(len(s))
	rest := n % int64(len(s))

	aCounter := func(s string, end int64) int64 {
		aCount := 0
		for i := 0; i < int(end); i++ {
			if s[i] == 'a' {
				aCount++
			}
		}
		return int64(aCount)
	}

	if !strings.Contains(s, "a") {
		return 0
	}
	if int64(len(s)) > n {
		return aCounter(s, rest)
	}
	return factor * aCounter(s, int64(len(s))) + aCounter(s, rest)
}

//func aCounter(s string, end int64) int64 {
//	aCount := 0
//	for i := 0; i < int(end); i++ {
//		if s[i] == 'a' {
//			aCount++
//		}
//	}
//	return int64(aCount)
//}