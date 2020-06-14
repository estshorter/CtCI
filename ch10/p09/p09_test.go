package p09

import "testing"

func Test_searchSortedMatrix(t *testing.T) {
	matrix := [][]int{{1, 2, 10}, {4, 6, 11}, {6, 7, 12}, {10, 11, 13}}
	type args struct {
		a [][]int
		x int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{matrix, 10}, 0, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := searchSortedMatrix(tt.args.a, tt.args.x)
			if got != tt.want {
				t.Errorf("searchSortedMatrix() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("searchSortedMatrix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}