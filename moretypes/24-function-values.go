package moretypes

import (
	"fmt"
	"math"
)

func Compute24(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FunctionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(Compute24(hypot))
	fmt.Println(Compute24(math.Pow))
}
