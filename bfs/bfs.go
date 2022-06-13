package bfs

import (
	"fmt"

	"github.com/PieroNarciso/bfs-go/graph"
	"github.com/PieroNarciso/bfs-go/vertex"
)

type BFS interface {
	Search(src vertex.Vertex, dest vertex.Vertex) int
	String() string
}

type bfs struct {
	graph  graph.Graph
	result int
}

func NewBFS(graph graph.Graph) BFS {
	return &bfs{graph: graph}
}

func (b *bfs) Search(src vertex.Vertex, dest vertex.Vertex) int {
	visited := map[vertex.Vertex]bool{}
	distance := map[vertex.Vertex]int{}
	queue := []vertex.Vertex{}
	queue = append(queue, src)
	visited[queue[0]] = true
	distance[queue[0]] = 0

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, neighboor := range b.graph.GetChildren(front) {
			if _, ok := visited[neighboor]; !ok {
				queue = append(queue, neighboor)
				visited[neighboor] = true
				distance[neighboor] = distance[front] + 1
			}
		}
	}
	b.result = distance[dest]
	return distance[dest]
}

func (b *bfs) String() string {
	return fmt.Sprintf("%d", b.result)
}
