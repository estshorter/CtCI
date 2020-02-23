package p01

import (
	"ctci/container/graph"
	"ctci/container/queue"
	"ctci/container/stack"
)

func bfs(i, j int, g *graph.Graph) bool {
	seen := make([]bool, g.NumVertexes())
	q := queue.New()
	q.Enqueue(i)
	seen[i] = true
	for !q.IsEmpty() {
		for _, edge := range g.OutEdges(q.Dequeue()) {
			if edge == j {
				return true
			} else if !seen[edge] {
				q.Enqueue(edge)
				seen[edge] = true
			}
		}
	}
	return false
}

func dfs(i, j int, g *graph.Graph) bool {
	seen := make([]bool, g.NumVertexes())
	s := stack.New()
	s.Push(i)
	seen[i] = true
	for !s.IsEmpty() {
		for _, edge := range g.OutEdges(s.Pop()) {
			if edge == j {
				return true
			} else if !seen[edge] {
				s.Push(edge)
				seen[edge] = true
			}
		}
	}
	return false
}

func dfsRecursive(i, j int, g *graph.Graph, seen []bool) bool {
	seen[i] = true
	for _, edge := range g.OutEdges(i) {
		if edge == j {
			return true
		} else if !seen[edge] {
			if dfsRecursive(edge, j, g, seen) {
				return true
			}
		}
	}
	return false
}

func dfs2(i, j int, g *graph.Graph) bool {
	seen := make([]bool, g.NumVertexes())
	return dfsRecursive(i, j, g, seen)
}
