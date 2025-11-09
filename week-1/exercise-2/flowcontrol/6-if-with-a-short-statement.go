package flowcontrol

import (
	"fmt"
	"math"
)

func Pow6(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func IfWithAShortStatement() {
	fmt.Println(
		Pow6(3, 2, 10),
		Pow6(3, 3, 10),
	)
}
