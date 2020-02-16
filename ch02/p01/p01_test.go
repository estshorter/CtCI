package p01

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func equal(l1 *list.List, l2 *list.List) bool {
	if l1.Len() != l2.Len() {
		fmt.Printf("wrong length: got %v, want %v\n", l1.Len(), l2.Len())
		return false
	}

	e1 := l1.Front()
	e2 := l2.Front()
	for i := 0; i < l1.Len(); i++ {
		if e1.Value != e2.Value {
			fmt.Printf("wrong value: got %v, want %v\n", e1.Value, e2.Value)
			return false
		}
		e1 = e1.Next()
		e2 = e2.Next()
	}
	return true
}

func TestRemoveDuplicate(t *testing.T) {
	l1 := list.New()
	l1.PushBack(3)
	l1.PushBack(5)
	l1.PushBack(1)
	l1.PushBack(100)
	l1.PushBack(1)
	l1.PushBack(10)
	l1.PushBack(5)
	l1.PushBack(10)

	l1Ret := list.New()
	l1Ret.PushBack(3)
	l1Ret.PushBack(5)
	l1Ret.PushBack(1)
	l1Ret.PushBack(100)
	l1Ret.PushBack(10)

	type args struct {
		l *list.List
	}

	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"1", args{l1}, l1Ret},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicate(tt.args.l)
			if !equal(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateLowMemory(t *testing.T) {
	l1 := list.New()
	l1.PushBack(3)
	l1.PushBack(5)
	l1.PushBack(1)
	l1.PushBack(100)
	l1.PushBack(1)
	l1.PushBack(10)
	l1.PushBack(5)
	l1.PushBack(10)

	l1Ret := list.New()
	l1Ret.PushBack(3)
	l1Ret.PushBack(5)
	l1Ret.PushBack(1)
	l1Ret.PushBack(100)
	l1Ret.PushBack(10)

	type args struct {
		l *list.List
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		{"1", args{l1}, l1Ret},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicateLowMemory(tt.args.l)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicateLowMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
