package interview

import (
	"regexp"
	"strings"
)

func IsPalindrome(text string) bool {
	text = sanitize(text)
	// more explicit process:
	// midIdx := len(text) / 2
	// lastIdx := len(text) - 1
	// for i := 0; i < midIdx; i++ {
	// 	if text[i] != text[lastIdx-i] {
	// 		return false
	// 	}
	// }

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func sanitize(text string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(text, "")
	return strings.ToLower(strings.Trim(safe, ""))
}

func recursivePalindrome(text string, forward int, backward int) bool {
	if forward == backward {
		return true
	}

	if text[forward] != text[backward] {
		return false
	}
	if forward < backward+1 {
		return recursivePalindrome(text, forward+1, backward-1)
	}
	return true
}

// Complete the gameOfThrones function below.
// check to see if an anagram is a palindrome.
// For example, given the string s = [aabbccdd], one way it can be arranged into a palindrome is abcddcba.
// if length is even then number of letter paris is even
// if length is odd then even number of

// maximum one element can have odd count, rest should be even.
// ex. aaabb -> baaab -> palindrome
func gameOfThrones(s string) string {

	s = sanitize(s)

	split := strings.Split(s, "")

	countMap := make(map[string]int)

	for _, element := range split {

		if _, ok := countMap[element]; !ok {
			count := strings.Count(s, element)
			countMap[element] = count
		}
	}

	flag := 0

	for _, value := range countMap {
		if value%2 != 0 {
			flag++
			if flag > 1 {
				return "NO"
			}
		}
	}

	return "YES"
}

func euclideanAlgorithm(a, b int32) int32 {
	if b == 0 {
		return b
	}
	remainder := a % b
	return euclideanAlgorithm(b, remainder)
}

func longestPalindrome(s string) string {
	return ""
}
