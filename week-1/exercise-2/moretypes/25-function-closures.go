package moretypes

import "fmt"

func Adder25() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func FunctionClosures() {
	pos, neg := Adder25(), Adder25()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
