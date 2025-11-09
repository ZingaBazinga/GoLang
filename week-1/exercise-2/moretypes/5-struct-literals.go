package moretypes

import "fmt"

type Vertex5 struct {
	X, Y int
}

var (
	v1 = Vertex5{1, 2}
	v2 = Vertex5{X: 1}
	v3 = Vertex5{1, 2}
	p  = &Vertex5{1, 2}
)

func StructLiterals() {
	fmt.Println(v1, p, v2, v3)
}
