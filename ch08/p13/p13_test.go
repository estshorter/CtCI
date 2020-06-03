package p13

import "testing"

func Test_createStack(t *testing.T) {
	type args struct {
		boxes []Box
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]Box{{2, 2, 2}, {1, 1, 1}, {3, 3, 3}, {1, 4, 4}}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createStack(tt.args.boxes); got != tt.want {
				t.Errorf("createStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
