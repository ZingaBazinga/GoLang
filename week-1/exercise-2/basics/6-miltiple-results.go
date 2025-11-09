package basics

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func MultipleResults() {
	a, b := swap("world", "hello")
	fmt.Println(a, b)
}