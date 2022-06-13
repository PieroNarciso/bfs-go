package graph

import (
	"github.com/PieroNarciso/bfs-go/edge"
	"github.com/PieroNarciso/bfs-go/vertex"
)

// Graph represents a graph
type Graph interface {
	AddEdge(src vertex.Vertex, dest vertex.Vertex)
	GetChildren(src vertex.Vertex) []vertex.Vertex
}

type dgraph struct {
	vertices map[vertex.Vertex][]edge.Edge
}

func NewDGraph() Graph {
	return &dgraph{
		vertices: map[vertex.Vertex][]edge.Edge{},
	}
}

func (g *dgraph) AddEdge(src vertex.Vertex, dest vertex.Vertex) {
	g.vertices[src] = append(g.vertices[src], edge.NewEdge(src, dest))
}

func (g *dgraph) GetChildren(src vertex.Vertex) []vertex.Vertex {
	edges, ok := g.vertices[src]
	if !ok {
		return []vertex.Vertex{}
	}
	var res []vertex.Vertex
	for _, edge := range edges {
		res = append(res, edge.Dest())
	}
	return res
}
