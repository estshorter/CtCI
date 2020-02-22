package p05

import "ctci/container/stack"

func sort(s *stack.Stack) *stack.Stack {
	s2 := stack.New()
	for !s.IsEmpty() {
		v := s.Pop()
		for !s2.IsEmpty() && v < s2.Peek() {
			s.Push(s2.Pop())
		}
		s2.Push(v)
	}
	for !s2.IsEmpty() {
		s.Push(s2.Pop())
	}
	return s
}
