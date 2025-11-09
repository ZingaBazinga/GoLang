package moretypes

import "fmt"

func MakingSlices() {
	a := make([]int, 5)
	PrintSlices13("a", a)

	b := make([]int, 0, 5)
	PrintSlices13("b", b)

	c := b[:2]
	PrintSlices13("c", c)

	d := c[2:5]
	PrintSlices13("d", d)
}

func PrintSlices13(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
