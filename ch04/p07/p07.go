package p07

import (
	"ctci/container/graph"
	"ctci/container/queue"
	"ctci/container/stack"
	"errors"
)

// FindOrder solves ch04-p07
func FindOrder(g *graph.Graph) ([]int, error) {
	var result []int
	NumDependency := make([]int, g.NumVertexes())
	q := queue.New()
	for vd := 0; vd < g.NumVertexes(); vd++ {
		lenInEdges := len(g.InEdges(vd))
		NumDependency[vd] = lenInEdges
		if lenInEdges == 0 {
			q.Enqueue(vd)
			result = append(result, vd)
		}
	}
	for !q.IsEmpty() {
		vs := q.Dequeue()
		for _, vd := range g.OutEdges(vs) {
			NumDependency[vd]--
			if NumDependency[vd] == 0 {
				q.Enqueue(vd)
				result = append(result, vd)
			}
		}
	}
	if len(result) != g.NumVertexes() {
		return nil, errors.New("No executable order")
	}
	return result, nil
}

type status int

const (
	blank status = iota
	partial
	completed
)

// FindOrderDFS solves ch04-p07
func FindOrderDFS(g *graph.Graph) ([]int, error) {
	s := stack.New()
	statusList := make([]status, g.NumVertexes())
	for i := 0; i < g.NumVertexes(); i++ {
		if statusList[i] == blank {
			if err := dfs(i, g, s, statusList); err != nil {
				return nil, err
			}
		}
	}
	ret := make([]int, s.Len())
	copy(ret, s.Slice())
	mid := len(ret) / 2
	for i := 0; i < mid; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret, nil
}

func dfs(i int, g *graph.Graph, s *stack.Stack, sl []status) error {
	sl[i] = partial
	for _, j := range g.OutEdges(i) {
		switch sl[j] {
		case blank:
			if err := dfs(j, g, s, sl); err != nil {
				return err
			}
		case partial:
			return errors.New("No executable order")
		}
	}
	s.Push(i)
	sl[i] = completed
	return nil
}
