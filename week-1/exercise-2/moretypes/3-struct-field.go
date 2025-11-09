package moretypes

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func StructField() {
	v := Vertex3{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}
