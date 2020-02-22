package p06

import (
	"ctci/ch03/p06/animal"
	"testing"
	"time"
)

func TestAnimalQueues(t *testing.T) {
	sleep := func() {
		time.Sleep(time.Nanosecond * 10)
	}

	q := New()
	kinds := []animal.Kind{animal.Cat, animal.Cat, animal.Dog,
		animal.Cat, animal.Cat, animal.Cat,
		animal.Dog}
	for _, kind := range kinds {
		q.Enqueue(kind)
		sleep()
	}
	got := q.DequeueAny().Kind
	want := animal.Cat
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
	got = q.DequeueCat().Kind
	want = animal.Cat
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
	got = q.DequeueDog().Kind
	want = animal.Dog
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}

}
