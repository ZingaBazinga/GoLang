package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt5(x float64) string {
	if x < 0 {
		return Sqrt5(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func If() {
	fmt.Println(Sqrt5(2), Sqrt5(-4))
}
