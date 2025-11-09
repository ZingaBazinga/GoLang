package moretypes

import "fmt"

func Fibonacci25() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func ExerciseFibonacciClosure() {
	f := Fibonacci25()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
