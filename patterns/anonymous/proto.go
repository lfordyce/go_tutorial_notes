package anonymous

type ProtoThing struct {
	itemMethod    func() float64
	setItemMethod func(float64)
}

func (t ProtoThing) Item() float64 {
	return t.itemMethod()
}

func (t ProtoThing) SetItem(x float64) {
	t.setItemMethod(x)
}

type Thing interface {
	Item() float64
	SetItem(float64)
}

func newThing() Thing {
	item := 0.0

	//var item float64
	//item = 0.0

	t := struct{ ProtoThing }{}

	t.itemMethod = func() float64 {
		return item
	}

	t.setItemMethod = func(x float64) {
		item = x
	}
	return t

	//return ProtoThing{
	//	itemMethod: func() float64 {
	//		return item
	//	},
	//	setItemMethod: func(x float64) {
	//		item = x
	//	},
	//}
}
