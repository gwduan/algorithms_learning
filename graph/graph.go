package graph

import (
	"github.com/gwduan/algorithms_learning/queue"
)

type element struct {
	v    int
	next *element
}

type Graph struct {
	v   int
	e   int
	adj []element
}

func NewGraph(v int) *Graph {
	return &Graph{
		v:   v,
		adj: make([]element, v),
	}
}

func (g *Graph) AddEdge(a int, b int) {
	if a < 0 || a > g.v-1 {
		return
	}
	if b < 0 || b > g.v-1 {
		return
	}

	e := &element{
		v:    b,
		next: g.adj[a].next,
	}
	g.adj[a].next = e
	g.adj[a].v++

	e = &element{
		v:    a,
		next: g.adj[b].next,
	}
	g.adj[b].next = e
	g.adj[b].v++

	g.e++
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) Degree(v int) int {
	return g.adj[v].v
}

func (g *Graph) MaxDegree() int {
	max := 0
	for i := 0; i < g.v; i++ {
		if g.adj[i].v > max {
			max = g.adj[i].v
		}
	}

	return max
}

func (g *Graph) DFS(v int) []int {
	if v < 0 || v > g.v-1 {
		return nil
	}

	marked := make([]bool, g.v)
	vs := make([]int, 0)
	vs = g.dfs(v, marked, vs)

	return vs
}

func (g *Graph) dfs(v int, marked []bool, vs []int) []int {
	marked[v] = true
	vs = append(vs, v)
	for p := g.adj[v].next; p != nil; p = p.next {
		if !marked[p.v] {
			vs = g.dfs(p.v, marked, vs)
		}
	}

	return vs
}

func (g *Graph) BFS(v int) []int {
	if v < 0 || v > g.v-1 {
		return nil
	}

	marked := make([]bool, g.v)
	vs := make([]int, 0)
	vs = g.bfs(v, marked, vs)

	return vs
}

func (g *Graph) bfs(v int, marked []bool, vs []int) []int {
	q := queue.NewListQueue()

	marked[v] = true
	vs = append(vs, v)
	q.Put(v)

	for !q.IsEmpty() {
		value, _ := q.Get()
		v = value.(int)
		for p := g.adj[v].next; p != nil; p = p.next {
			if !marked[p.v] {
				marked[p.v] = true
				vs = append(vs, p.v)
				q.Put(p.v)
			}
		}
	}

	return vs
}
