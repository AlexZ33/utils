package containers

import "errors"

// Stack 栈
type Stack struct {
	top  int
	data []interface{}
}

// NewStack 创建栈
func NewStack() *Stack {
	return &Stack{
		top:  -1,
		data: make([]interface{}, 0, 32),
	}
}

// Push 入栈
func (s *Stack) Push(v interface{}) {
	s.top++
	s.data = append(s.data, v)
}

// Pop 出栈
func (s *Stack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("stack is empty")
	}
	v := s.data[s.top]
	s.top--
	s.data = s.data[:s.top+1]
	return v, nil
}

// Top 返回栈顶元素
func (s *Stack) Top() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("stack is empty")
	}
	return s.data[s.top], nil
}

// IsEmpty 栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

// Size 栈的大小
func (s *Stack) Size() int {
	return s.top + 1
}
