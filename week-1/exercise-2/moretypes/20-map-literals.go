package moretypes

import "fmt"

type Vertex20 struct {
	Lat, Long float64
}

var m20 = map[string]Vertex20{
	"Bell Labs": Vertex20{
		40.68433, -74.39967,
	},
	"Google": Vertex20{
		37.42202, -122.08408,
	},
}

func MapLiterals() {
	fmt.Println(m20)
}
