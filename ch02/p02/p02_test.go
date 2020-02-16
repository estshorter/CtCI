package p02

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_findLastKElement(t *testing.T) {
	l1 := list.New()
	l1.PushBack(3)
	l1.PushBack(5)
	e3 := l1.PushBack(1)
	l1.PushBack(100)
	e5 := l1.PushBack(10)

	type args struct {
		l *list.List
		k int
	}
	tests := []struct {
		name    string
		args    args
		want    *list.Element
		wantErr bool
	}{
		{"0", args{l1, 0}, e5, false},
		{"2", args{l1, 2}, e3, false},
		{"error", args{l1, 20}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLastKElement(tt.args.l, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("findLastKElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLastKElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
