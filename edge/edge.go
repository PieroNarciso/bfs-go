package edge

import "github.com/PieroNarciso/bfs-go/vertex"

type Edge interface {
	Src() vertex.Vertex
	Dest() vertex.Vertex
}

type edge struct {
	src    vertex.Vertex
	dest   vertex.Vertex
}

func NewEdge(src vertex.Vertex, dest vertex.Vertex) Edge {
	return &edge{
		src, dest,
	}
}

func (e *edge) Src() vertex.Vertex {
	return e.src
}

func (e *edge) Dest() vertex.Vertex {
	return e.dest
}
