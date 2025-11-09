package basics

import (
	"fmt"
	"math"
)

func ExportedNames() {
	// не fmt.Println(math.pi) должно быть с большой буквы Pi
	fmt.Println(math.Pi)
}