package json_utils

type RuneTypeCounts struct {
	vowels, constants, other int
}

func AnalyseText(text string) RuneTypeCounts {
	var result RuneTypeCounts
	m := make(map[rune]*int)
	for _, r := range "aeiou" {
		m[r] = &result.vowels
	}

	for _, r := range "bcdfghjklmnpqrstvwxyz" {
		m[r] = &result.constants
	}

	for _, r := range text {
		ref, ok := m[r]
		if ok {
			*ref++
		} else {
			result.other++
		}
	}
	return result
}
