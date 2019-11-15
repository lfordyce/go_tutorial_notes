package utils

import "strings"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (day Weekday) String() string {
	// Declare an array of strings
	// ... operator counts how many
	// items in the array (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	// → `day`: It's one of the
	// values of Weekday constants.
	// If the constant is Sunday,
	// then day is 0.

	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < Sunday || day > Saturday {
		return "Unkown"
	}

	// return the name of a Weekday
	// constant from the names array
	// above.
	return names[day]
}

func (day Weekday) Weekend() bool {
	switch day {
	// If day is a weekend day:
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}

func customSplitter(s string, splits string) []string {
	m := make(map[rune]int)
	for _, r := range splits {
		m[r] = 1
	}

	splitter := func(r rune) bool {
		return m[r] == 1
	}

	return strings.FieldsFunc(s, splitter)
}
