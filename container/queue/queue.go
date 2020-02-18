package queue

// Queue implements queue data structure
type Queue struct {
	a []int
	j int
	n int
}

// New initializes queue
func New() *Queue {
	return &Queue{a: []int{}}
}

// Len returns number of elements
func (q *Queue) Len() int {
	return q.n
}

// Enqueue set a value v to the last of Queue
func (q *Queue) Enqueue(v int) {
	if q.n+1 > cap(q.a) {
		q.resize()
	}
	q.a[(q.j+q.n)%cap(q.a)] = v
	q.n++
}

// Dequeue gets the last element v of a Queue and delete it from the Queue
func (q *Queue) Dequeue() int {
	v := q.a[q.j]
	q.j = (q.j + 1) % cap(q.a)
	q.n--
	if cap(q.a) >= 3*q.n {
		q.resize()
	}
	return v
}

func (q *Queue) resize() {
	b := make([]int, max(2*q.n, 1))
	for k := 0; k < q.n; k++ {
		b[k] = q.a[(q.j+k)%cap(q.a)]
	}
	q.a = b
	q.j = 0
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
