package vertex_test

import (
	"testing"

	. "github.com/PieroNarciso/bfs-go/vertex"
)

func TestVextex(t *testing.T) {
	v := NewVertex("a")
	if val := v.Name(); val != "a" {
		t.Error("Error value in Value() with", "a")
	}

	v = NewVertex("b")
	if val := v.Name(); val != "b" {
		t.Error("Error value in Value() with", "b")
	}
}
