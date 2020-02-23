package p01

import (
	"ctci/container/graph"
	"testing"
)

func Test_bfs(t *testing.T) {
	g := graph.New(5, [][]int{{0, 1}, {0, 4}, {1, 2}, {2, 1}, {2, 3}, {3, 2}, {4, 0}})
	type args struct {
		i int
		j int
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{4, 3, g}, true},
		{"case2", args{3, 0, g}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.i, tt.args.j, tt.args.g); got != tt.want {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfs(t *testing.T) {
	g := graph.New(5, [][]int{{0, 1}, {0, 4}, {1, 2}, {2, 1}, {2, 3}, {3, 2}, {4, 0}})
	type args struct {
		i int
		j int
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{4, 3, g}, true},
		{"case2", args{3, 0, g}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.i, tt.args.j, tt.args.g); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfs2(t *testing.T) {
	g := graph.New(5, [][]int{{0, 1}, {0, 4}, {1, 2}, {2, 1}, {2, 3}, {3, 2}, {4, 0}})
	type args struct {
		i int
		j int
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{4, 3, g}, true},
		{"case2", args{3, 0, g}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs2(tt.args.i, tt.args.j, tt.args.g); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
