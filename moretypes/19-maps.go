package moretypes

import "fmt"

type Vertex19 struct {
	Lat, Long float64
}

var m19 map[string]Vertex19

func Map() {
	m19 = make(map[string]Vertex19)

	m19["Bell Labs"] = Vertex19{
		40.68433, -74.39967,
	}
	fmt.Println(m19["Bell Labs"])
}
