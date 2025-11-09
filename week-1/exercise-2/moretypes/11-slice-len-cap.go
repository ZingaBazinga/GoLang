package moretypes

import "fmt"

func SliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	PrintSlices11(s)

	// Slice the slice to give it zero length.
	s = s[:0]

	// Extend its length.
	s = s[:4]
	PrintSlices11(s)

	// Drop its first two values.
	s = s[2:]
	PrintSlices11(s)
}

func PrintSlices11(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
