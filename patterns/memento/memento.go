package memento

type memento struct {
	money int
}

func (m *memento) getMoney() int {
	return m.money
}

type Game struct {
	Money int
}

func (g *Game) CreateMemento() *memento {
	return &memento{money: g.Money}
}

func (g *Game) RestoreMemento(memento *memento) {
	g.Money = memento.getMoney()
}
