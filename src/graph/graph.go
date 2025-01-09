package graph

import (
	"fmt"
)

type Graph struct {
	V   int
	Adj map[string][]Edge
}

func NewGraph(v int) *Graph {
	return &Graph{V: v, Adj: make(map[string][]Edge, v)}
}

func (g *Graph) AddEdge(v string, w string, weight int) {
	if g == nil {
		fmt.Println("Graph is nil!")
		return
	}

	if v == "" || w == "" {
		fmt.Println("Invalid edge: source or destination is empty")
		return
	}

	for _, edge := range g.Adj[v] {
		if edge.dest == w && edge.weight == weight {
			return // Ребро уже существует, выходим
		}
	}

	edge := Edge{src: v, dest: w, weight: weight}
	g.Adj[v] = append(g.Adj[v], edge)

	if _, exists := g.Adj[w]; !exists {
		g.Adj[w] = []Edge{}
	}

	edge = Edge{src: w, dest: v, weight: weight}
	g.Adj[w] = append(g.Adj[w], edge)
}

func (g *Graph) DFS() {
	visited := make(map[string]bool)
	for k := range g.Adj {
		visited[k] = false
	}
	for k := range g.Adj {
		if !visited[k] {
			g.DFSUtil(k, visited)
		}
	}
}

func (g *Graph) DFSUtil(k string, visited map[string]bool) {
	visited[k] = true
	fmt.Print(k + " ")
	for _, edge := range g.Adj[k] {
		if !visited[edge.dest] {
			g.DFSUtil(edge.dest, visited)
		}
	}
}

func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	var queue []string
	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Print(vertex + " ")
		for _, edge := range g.Adj[vertex] {
			if !visited[edge.dest] {
				visited[edge.dest] = true
				queue = append(queue, edge.dest)
			}
		}
	}
}

func (g *Graph) Print() {
	for vertex := range g.Adj {
		fmt.Printf("%s -> ", vertex)
		if len(g.Adj[vertex]) == 0 {
			fmt.Print("(no edges)")
		} else {
			// Иначе, выводим рёбра для этой вершины
			for _, edge := range g.Adj[vertex] {
				fmt.Printf("(%s, %d) ", edge.dest, edge.weight)
			}
		}
		fmt.Println()
	}

}

func (g *Graph) FindMinSpanningTree() {
	edges := make([]Edge, 0)
	averageWeight := 0
	for _, e := range g.Adj {
		for j := 0; j < len(e); j++ {
			edges = append(edges, e[j])
		}
	}
	edges = KruskalsAlgorithm(edges, g.V)
	fmt.Println("Edges in MST are:")
	edges = insertionSort(edges, false)
	for _, e := range edges {
		fmt.Printf("%s -- %s (%d)\n", e.src, e.dest, e.weight)
		averageWeight += e.weight
	}
	fmt.Printf("Average Weight = %d\n", averageWeight)
}

type Edge struct {
	src, dest string
	weight    int
}

func compareEdges(edge1, edge2 Edge, weight bool) bool {
	if weight {
		return edge1.weight > edge2.weight
	} else {
		if edge1.src != edge2.src {
			return edge1.src > edge2.src
		}
		return edge1.dest > edge2.dest
	}
}

func insertionSort(arr []Edge, weightSort bool) []Edge {
	var i, j int
	var key Edge
	for i = 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1
		for j >= 0 && compareEdges(arr[j], key, weightSort) {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func KruskalsAlgorithm(edges []Edge, numVertices int) []Edge {
	edges = insertionSort(edges, true)
	parent := make(map[string]string, numVertices)
	var minSpanningTree []Edge
	count := 0
	for _, e := range edges {
		srcParent := ""
		destParent := ""
		src := e.src
		dest := e.dest
		for parent[src] != "" {
			src = parent[src]
		}
		for parent[dest] != "" {
			dest = parent[dest]
		}
		srcParent = src
		destParent = dest
		if srcParent != destParent {
			minSpanningTree = append(minSpanningTree, e)
			count++
			parent[destParent] = srcParent
		}
		if count == numVertices-1 {
			break
		}
	}
	return minSpanningTree
}
