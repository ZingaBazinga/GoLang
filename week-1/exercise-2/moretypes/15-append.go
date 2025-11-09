package moretypes

import "fmt"

func Append() {
	var s []int
	PrintSlices15(s)

	s = append(s, 0)
	PrintSlices15(s)

	s = append(s, 1)
	PrintSlices15(s)

	s = append(s, 2, 3, 4)
	PrintSlices15(s)
}

func PrintSlices15(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
