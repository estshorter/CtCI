package sllist

import "errors"

//Element is an element of a linked list.
type Element struct {
	// The value stored with this element.
	Value int
	Next  *Element
}

//List represents a singly linked list.
type List struct {
	head *Element
	tail *Element
	N    int
}

//New returns an initialized list.
func New() *List {
	return &List{}
}

//InitBySlice returns an initialized list.
func InitBySlice(s []int) *List {
	l := &List{}
	l.PushBackSlice(s)
	return l
}

//IsConsumed returns true if e is dummy element
func (l *List) IsConsumed(e *Element) bool {
	return e == nil
}

//PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v int) *Element {
	e := &Element{v, l.head}
	l.head = e
	if l.N == 0 {
		l.tail = e
	}
	l.N++
	return e
}

//PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v int) *Element {
	return l.PushBackElement(&Element{v, nil})
}

//PushBackLoop inserts a new element e pointing the element whose number equals eNum
func (l *List) PushBackLoop(v int, eNum int) *Element {
	e := &Element{v, l.GetElement(eNum)}
	if l.N == 0 {
		panic("size is zero")
	} else {
		l.tail.Next = e
	}
	l.tail = e
	l.N++
	return e
}

//PushBackElement inserts an existing element at the back of list l and returns e.
func (l *List) PushBackElement(e *Element) *Element {
	e.Next = nil
	if l.N == 0 {
		l.head = e
	} else {
		l.tail.Next = e
	}
	l.tail = e
	l.N++
	return e
}

//PushBackSlice inserts a slice at the back of list l.
func (l *List) PushBackSlice(s []int) {
	for _, v := range s {
		l.PushBack(v)
	}
}

//GetElement returns ith element
func (l *List) GetElement(i int) *Element {
	if i < 0 || i >= l.N {
		return nil
	}
	var e *Element
	if i == l.N-1 {
		e = l.tail
	} else {
		e = l.head
		for j := 0; j < i; j++ {
			e = e.Next
		}
	}
	return e
}

//Front returns the first element of list.
func (l *List) Front() *Element {
	return l.head
}

// Back returns the last element of list l.
func (l *List) Back() *Element {
	return l.tail
}

//Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() int {
	return l.N
}

//RemoveFront removes a first element from l It returns the value of the removed element e.
func (l *List) RemoveFront() (int, error) {
	if l.N == 0 {
		return 0, errors.New("list is empty")
	}
	v := l.head.Value
	l.head = l.head.Next
	l.N--
	if l.N == 0 {
		l.tail = nil
	}
	return v, nil
}
