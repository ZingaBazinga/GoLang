package moretypes

import "fmt"

type Vertex4 struct {
	X int
	Y int
}

func StructPointers() {
	v := Vertex4{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
