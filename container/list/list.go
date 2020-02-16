package list

//Element is an element of a linked list.
type Element struct {
	// The value stored with this element.
	Value int
	Prev  *Element
	Next  *Element
}

//List represents a doubly linked list.
type List struct {
	dummy *Element
	N     int
}

//New returns an initialized list.
func New() *List {
	dummy := &Element{}
	dummy.Prev = dummy
	dummy.Next = dummy
	return &List{dummy: dummy}
}

//PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v int) *Element {
	e := l.InsertBefore(v, l.getElement(l.N+1))
	return e
}

func (l *List) getElement(i int) *Element {
	var e *Element
	if i < l.N/2 {
		e = l.dummy.Next
		for j := 0; j < i; j++ {
			e = e.Next
		}
	} else {
		e = l.dummy
		for j := l.N; j > i; j-- {
			e = e.Prev
		}
	}
	return e
}

//InsertBefore inserts a new element e with value v immediately before mark and returns e.
//If mark is not an element of l, the list is not modified. The mark must not be nil.
func (l *List) InsertBefore(v int, mark *Element) *Element {
	e := &Element{Value: v, Prev: mark.Prev, Next: mark}
	e.Next.Prev = e
	e.Prev.Next = e
	l.N++
	return e
}

//Front returns the first element of list.
func (l *List) Front() *Element {
	if l.N > 0 {
		return l.getElement(1)
	}
	return l.dummy
}

// Back returns the last element of list l.
func (l *List) Back() *Element {
	if l.N > 0 {
		return l.getElement(l.Len())
	}
	return l.dummy
}

// Dummy returns dummy element
func (l *List) Dummy() *Element {
	return l.dummy
}

//Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() int {
	return l.N
}

//Remove removes e from l if e is an element of list l. It returns the element value e.Value. The element must not be nil.
func (l *List) Remove(e *Element) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil
	e.Prev = nil
	l.N--
}
