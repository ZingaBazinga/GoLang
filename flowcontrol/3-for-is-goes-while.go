package flowcontrol

import "fmt"

func ForIsGoesWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
