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

//NewWithSlice returns an initialized list.
func NewWithSlice(s []int) *List {
	dummy := &Element{}
	dummy.Prev = dummy
	dummy.Next = dummy
	l := &List{dummy: dummy}
	l.PushBackSlice(s)
	return l
}

//PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v int) *Element {
	e := l.InsertBefore(v, l.GetElement(l.N+1))
	return e
}

//PushBackElement inserts an existing element at the back of list l and returns e.
func (l *List) PushBackElement(e *Element) *Element {
	l.InsertBeforeElement(e, l.GetElement(l.N+1))
	return e
}

//PushBackList inserts an existing List l2 at the back of list l.
//l2 must not be used after this.
func (l *List) PushBackList(l2 *List) {
	l2.Dummy().Next.Prev = l.Dummy().Prev
	l.Dummy().Prev.Next = l2.Dummy().Next
	l2.Dummy().Prev.Next = l.Dummy()
	l.Dummy().Prev = l2.Dummy().Prev
	l.N += l2.N
}

//PushBackSlice inserts a slice at the back of list l.
func (l *List) PushBackSlice(s []int) {
	for _, v := range s {
		l.PushBack(v)
	}
}

//GetElement returns ith element
func (l *List) GetElement(i int) *Element {
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

//InsertBeforeElement InsertBefore inserts an existing element e immediately before mark and returns e.
func (l *List) InsertBeforeElement(e *Element, mark *Element) *Element {
	e.Prev = mark.Prev
	e.Next = mark
	e.Next.Prev = e
	e.Prev.Next = e
	l.N++
	return e
}

//Front returns the first element of list.
func (l *List) Front() *Element {
	return l.GetElement(0)
}

// Back returns the last element of list l.
func (l *List) Back() *Element {
	return l.GetElement(l.Len() - 1)
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
