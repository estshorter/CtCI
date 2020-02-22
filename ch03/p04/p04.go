package p04

import "ctci/container/stack"

// MyQueue implements queue interface by using 2 stacks
type MyQueue struct {
	newest stack.Stack
	oldest stack.Stack
}

// New initializes MyQueue
func New() *MyQueue {
	return &MyQueue{*stack.New(), *stack.New()}
}

// Enqueue adds data v to MyQueue
func (q *MyQueue) Enqueue(v int) {
	q.newest.Push(v)
}

// Dequeue removes data  from MyQueue and returns it
func (q *MyQueue) Dequeue() int {
	if q.oldest.IsEmpty() {
		for !q.newest.IsEmpty() {
			q.oldest.Push(q.newest.Pop())
		}
	}
	return q.oldest.Pop()
}
