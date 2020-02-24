package p07

import (
	"ctci/container/graph"
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	edges := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}}
	g := graph.New(6, edges)

	edges2 := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}, {2, 4}, {4, 5}}
	g2 := graph.New(6, edges2)

	edges3 := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}, {2, 4}}
	g3 := graph.New(6, edges3)

	type args struct {
		g *graph.Graph
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"case1", args{g}, []int{4, 5, 1, 0, 3, 2}, false},
		{"case2", args{g2}, nil, true},
		{"case3", args{g3}, []int{5, 1, 0, 3, 2, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindOrder(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOrderDFS(t *testing.T) {
	edges := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}}
	g := graph.New(6, edges)

	edges2 := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}, {2, 4}, {4, 5}}
	g2 := graph.New(6, edges2)

	edges3 := [][]int{{0, 3}, {1, 0}, {1, 3}, {3, 2}, {5, 0}, {5, 1}, {2, 4}}
	g3 := graph.New(6, edges3)

	type args struct {
		g *graph.Graph
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"case1", args{g}, []int{5, 4, 1, 0, 3, 2}, false},
		{"case2", args{g2}, nil, true},
		{"case3", args{g3}, []int{5, 1, 0, 3, 2, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindOrderDFS(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrderDFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrderDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
