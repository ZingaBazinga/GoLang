package moretypes

// import "golang.org/x/tour/pic"
// Запускается только на https://go.dev/tour/moretypes/18

func Pic18(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			result[y][x] = uint8((x + y) / 2)
		}
	}
	return result
}

func ExerciseSlices() {
	// Pic18.Show(Pic)
}
