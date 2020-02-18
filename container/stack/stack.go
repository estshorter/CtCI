package stack

// Stack implements stack data structure
type Stack struct {
	a []int
}

// New initializes stack
func New() *Stack {
	return &Stack{[]int{}}
}

// Len returns number of elements
func (s *Stack) Len() int {
	return len(s.a)
}

// Push pushes a value v to a stack
func (s *Stack) Push(v int) {
	// s.a = append(s.a, v)
	if len(s.a)+1 > cap(s.a) {
		s.resize()
	}
	s.a = s.a[:len(s.a)+1]
	s.a[len(s.a)-1] = v
}

// Pop pops a value v from a stack
func (s *Stack) Pop() int {
	v := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]
	if cap(s.a) >= 3*len(s.a) {
		s.resize()
	}
	return v
}

func (s *Stack) resize() {
	b := make([]int, len(s.a), max(2*len(s.a), 1))
	copy(b, s.a)
	s.a = b
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
