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
		k := q.Dequeue()
		edges := g.OutEdges(k)
		for _, edge := range edges {
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
		r := s.Pop()
		edges := g.OutEdges(r)
		for _, edge := range edges {
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
