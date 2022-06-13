package main

import (
	"fmt"

	"github.com/PieroNarciso/bfs-go/bfs"
	"github.com/PieroNarciso/bfs-go/graph"
	"github.com/PieroNarciso/bfs-go/vertex"
)

func main() {
	graph := graph.NewDGraph()

	va := vertex.NewVertex("A")
	vb := vertex.NewVertex("B")
	vc := vertex.NewVertex("C")
	vd := vertex.NewVertex("D")
	ve := vertex.NewVertex("E")

	graph.AddEdge(va, vb)
	graph.AddEdge(vb, vc)
	graph.AddEdge(vc, vd)
	graph.AddEdge(vd, ve)

	graph.AddEdge(va, vd)

	bfs := bfs.NewBFS(graph)
	bfs.Search(va, ve)
	fmt.Println(bfs.String())
}
