package vertex

type Vertex interface {
	Name() string
}

type vertex struct {
	name string
}

func NewVertex(name string) Vertex {
	return &vertex{name: name}
}

func (v *vertex) Name() string {
	return v.name
}
