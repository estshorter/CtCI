package p02

import "ctci/container/stack"

const maxInt = int(^uint(0) >> 1)

// StackWithMin impelemnts a stack structure which can return minimize value
type StackWithMin struct {
	s  stack.Stack
	sm stack.Stack
}

// New initializes StackWithMin
func New() *StackWithMin {
	return &StackWithMin{*stack.New(), *stack.New()}
}

// Push adds a val to stack
func (s *StackWithMin) Push(val int) {
	if s.sm.Len() == 0 || val <= s.Min() {
		s.sm.Push(val)
	}
	s.s.Push(val)
}

// Pop removes a value from stack and returns it
func (s *StackWithMin) Pop() int {
	v := s.s.Pop()
	if v == s.Min() {
		s.sm.Pop()
	}
	return v
}

// Min returns the minimum value in stack
func (s *StackWithMin) Min() int {
	return s.sm.Peek()
}
