package flowcontrol

import (
	"fmt"
	"math"
)

func Sqrt8(x float64) float64 {
	if x < 0 {
		// Квадратный корень из отрицательного числа — не определён в ℝ
		return math.NaN()
	}
	if x == 0 {
		return 0
	}

	z := x
	const epsilon = 1e-10

	for {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(oldZ-z) < epsilon {
			break
		}
	}

	return z
}

func ExerviseLoopsAndFunctions() {
	fmt.Println(Sqrt8(2))
	fmt.Println(math.Sqrt(2))
}
