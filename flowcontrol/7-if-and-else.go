package flowcontrol

import (
	"fmt"
	"math"
)

func Pow7(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func IfAndElse() {
	fmt.Println(
		Pow7(3, 2, 10),
		Pow7(3, 3, 20),
	)
}
