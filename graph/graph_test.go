package graph

import (
	"fmt"
	"slices"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph(13)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 5)
	g.AddEdge(0, 6)

	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	g.AddEdge(4, 5)
	g.AddEdge(4, 6)

	g.AddEdge(7, 8)

	g.AddEdge(9, 10)
	g.AddEdge(9, 11)
	g.AddEdge(9, 12)

	g.AddEdge(11, 12)

	if got := g.V(); got != 13 {
		t.Errorf("V() = %v, want %v", got, 13)
	}
	if got := g.E(); got != 13 {
		t.Errorf("E() = %v, want %v", got, 13)
	}

	if got := g.Degree(0); got != 4 {
		t.Errorf("Degree(0) = %v, want %v", got, 4)
	}
	if got := g.Degree(6); got != 2 {
		t.Errorf("Degree(6) = %v, want %v", got, 2)
	}
	if got := g.Degree(9); got != 3 {
		t.Errorf("Degree(9) = %v, want %v", got, 3)
	}
	if got := g.MaxDegree(); got != 4 {
		t.Errorf("MaxDegree() = %v, want %v", got, 4)
	}

	printGraph(g)

	vs := g.DFS(0)
	//fmt.Printf("DFS v%-2d[%d]: %v\n", 0, len(vs), vs)
	wants := []int{0, 6, 4, 5, 3, 2, 1}
	if !slices.Equal(vs, wants) {
		t.Errorf("DFS(0) = %v, wants %v\n", vs, wants)
	}
	vs = g.DFS(3)
	//fmt.Printf("DFS v%-2d[%d]: %v\n", 3, len(vs), vs)
	wants = []int{3, 5, 4, 6, 0, 2, 1}
	if !slices.Equal(vs, wants) {
		t.Errorf("DFS(3) = %v, wants %v\n", vs, wants)
	}
	vs = g.DFS(8)
	//fmt.Printf("DFS v%-2d[%d]: %v\n", 8, len(vs), vs)
	wants = []int{8, 7}
	if !slices.Equal(vs, wants) {
		t.Errorf("DFS(8) = %v, wants %v\n", vs, wants)
	}
	vs = g.DFS(10)
	//fmt.Printf("DFS v%-2d[%d]: %v\n", 10, len(vs), vs)
	wants = []int{10, 9, 12, 11}
	if !slices.Equal(vs, wants) {
		t.Errorf("DFS(10) = %v, wants %v\n", vs, wants)
	}

	vs = g.BFS(0)
	//fmt.Printf("BFS v%-2d[%d]: %v\n", 0, len(vs), vs)
	wants = []int{0, 6, 5, 2, 1, 4, 3}
	if !slices.Equal(vs, wants) {
		t.Errorf("BFS(0) = %v, wants %v\n", vs, wants)
	}
	vs = g.BFS(3)
	//fmt.Printf("BFS v%-2d[%d]: %v\n", 3, len(vs), vs)
	wants = []int{3, 5, 4, 0, 6, 2, 1}
	if !slices.Equal(vs, wants) {
		t.Errorf("BFS(3) = %v, wants %v\n", vs, wants)
	}
	vs = g.BFS(8)
	//fmt.Printf("BFS v%-2d[%d]: %v\n", 8, len(vs), vs)
	wants = []int{8, 7}
	if !slices.Equal(vs, wants) {
		t.Errorf("BFS(8) = %v, wants %v\n", vs, wants)
	}
	vs = g.BFS(10)
	//fmt.Printf("BFS v%-2d[%d]: %v\n", 10, len(vs), vs)
	wants = []int{10, 9, 12, 11}
	if !slices.Equal(vs, wants) {
		t.Errorf("BFS(10) = %v, wants %v\n", vs, wants)
	}
}

func printGraph(g *Graph) {
	for i := 0; i < g.v; i++ {
		fmt.Printf("v%-2d[%d]: ", i, g.adj[i].v)
		for j := g.adj[i].next; j != nil; j = j.next {
			fmt.Printf("%-2d ", j.v)
		}
		fmt.Println()
	}
}
