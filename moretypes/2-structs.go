package moretypes

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func Structs() {
	fmt.Println(Vertex2{1, 2})
}
