package strategy

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Strategy func(int, int) int

type Strategic interface {
	SetStrategy(Strategy)
	Result() int
}

type secretStrategy struct {
	a        int
	b        int
	result   int
	strategy Strategy
}

func (ss *secretStrategy) SetStrategy(s Strategy) {
	ss.strategy = s
}

func (ss *secretStrategy) Result() int {
	ss.result = ss.strategy(ss.a, ss.b)
	return ss.result
}

func New(a, b int) Strategic {
	return &secretStrategy{a: a, b: b}
}
