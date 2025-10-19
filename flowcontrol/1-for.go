package flowcontrol

import (
	"fmt"
)

func For() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
