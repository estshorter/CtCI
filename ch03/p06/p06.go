package p06

import (
	"ctci/ch03/p06/animal"
	"ctci/ch03/p06/queue"
	"fmt"
	"time"
)

// AnimalQueue implements ch02-p06
type AnimalQueue struct {
	cat queue.Queue
	dog queue.Queue
}

// Enqueue add animal dat
func (a *AnimalQueue) Enqueue(kind animal.Kind) {
	switch kind {
	case animal.Cat:
		a.cat.Enqueue(animal.Animal{Kind: animal.Cat, Time: time.Now()})
	case animal.Dog:
		a.dog.Enqueue(animal.Animal{Kind: animal.Dog, Time: time.Now()})
	default:
	}
}

// DequeueAny dequeues the latest animal data
func (a *AnimalQueue) DequeueAny() animal.Animal {
	catTime := a.cat.Peek().Time
	dogTime := a.dog.Peek().Time
	fmt.Printf("%v, %v", catTime, dogTime)
	if catTime.Before(dogTime) {
		return a.cat.Dequeue()
	}
	return a.dog.Dequeue()
}

// DequeueDog dequeues data from the animal.Dog queue
func (a *AnimalQueue) DequeueDog() animal.Animal {
	return a.dog.Dequeue()
}

// DequeueCat dequeues data from the cat queue
func (a *AnimalQueue) DequeueCat() animal.Animal {
	return a.cat.Dequeue()
}

// New initializes AnimalQueue
func New() *AnimalQueue {
	return &AnimalQueue{*queue.New(), *queue.New()}
}
