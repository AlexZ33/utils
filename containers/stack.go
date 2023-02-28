package containers

type xStackElement struct {
	next  *xStackElement
	stack *xStack
	Value interface{}
}

type xStack struct {
	root xStackElement
	len  int
}

func (s *xStack) init() *xStack {
	s.root.next = &s.root
	s.len = 0
	return s
}

func NewxStack() *xStack {
	return new(xStack).init()
}

func (s *xStack) Push(v interface{}) {
	se := xStackElement{s.root.next}
}
