package interview

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}

func intToRomanRev(num int) string {
	values := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbols := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	res, i := "", len(symbols)-1
	for num != 0 {
		for values[i] <= num {
			num -= values[i]
			res += symbols[i]
		}
		i--
	}
	return res
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// MCMXCIV -> 1994
// "IV" -> 4
// "IX" -> 9
// "VII" -> 7
func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}
