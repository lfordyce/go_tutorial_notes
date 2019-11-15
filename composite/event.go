package composite

type Stats struct {
	cnt        int
	categories map[string]*Events
}

func (s *Stats) Category(n string) (e *Events) {
	if s.categories == nil {
		s.categories = map[string]*Events{}
	}
	if e = s.categories[n]; e == nil {
		e = &Events{}
		s.categories[n] = e
	}
	return
}

type Events struct {
	cnt    int
	events map[string]*Event
}

func (e *Events) Event(n string) (ev *Event) {
	if e.events == nil {
		e.events = map[string]*Event{}
	}
	if ev = e.events[n]; ev == nil {
		ev = &Event{}
		e.events[n] = ev
	}
	return
}

type Event struct {
	value int64
}

func romanNumeralDict() func(int) string {
	// innerMap is captured in the closure returned below
	innerMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	return func(key int) string {
		return innerMap[key]
	}
}
