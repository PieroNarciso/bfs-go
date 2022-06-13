package edge_test

import (
	"testing"

	"github.com/PieroNarciso/bfs-go/edge"
	"github.com/PieroNarciso/bfs-go/vertex"
)

func TestEdge(t *testing.T) {
	src := vertex.NewVertex("a")
	dest := vertex.NewVertex("b")
	edge := edge.NewEdge(src, dest)

	if srcVal := edge.Src().Name(); srcVal != "a" {
		t.Error("Error in value src Name() with", "a")
	}

	if destVal := edge.Dest().Name(); destVal != "b" {
		t.Error("Error in value src Name() with", "b")
	}
}
