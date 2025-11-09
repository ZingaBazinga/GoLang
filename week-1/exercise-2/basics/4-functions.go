package basics

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func Functions() {
	fmt.Println(Add(42, 13))
}