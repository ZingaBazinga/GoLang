//--------------------------------------------//
// basics/1
// packages.go
// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is", rand.Intn(100))
// }

//--------------------------------------------//
// basics/2
// imports.go
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
// }

//--------------------------------------------//
// basics/3
// exported-names.go
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// не fmt.Println(math.pi) должно быть с большой буквы Pi
// 	fmt.Println(math.Pi)
// }

// --------------------------------------------//
// basics/4
// functions.go
// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

// --------------------------------------------//
// basics/5
// functions-continued.go
// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

// --------------------------------------------//
// basics/6
// miltiple-results.go
// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("world", "hello")
// 	fmt.Println(a, b)
// }

// --------------------------------------------//
// basics/7
// named-results.go
// package main

// import "fmt"

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// func main() {
// 	fmt.Println(split(17))
// }

// --------------------------------------------//
// basics/8
// variables.go
// package main

// import "fmt"

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

// --------------------------------------------//
// basics/9
// variables-with-initializers.go
// package main

// import "fmt"

// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

// --------------------------------------------//
// basics/10
// short-variable-declarations.go
// package main

// import "fmt"

// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no!"
// 	fmt.Println(i, j, k, c, python, java)
// }

// --------------------------------------------//
// basics/11
// basic-types.go
// package main

// import (
// 	"fmt"
// 	"math/cmplx"
// )

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func main() {
// 	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
// 	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
// 	fmt.Printf("Type %T Value: %v\n", z, z)
// }

// --------------------------------------------//
// basics/12
// zero.go
// package main

// import "fmt"

// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// --------------------------------------------//
// basics/13
// type-conversions.go
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var x, y int = 3, 4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))
// 	var z uint = uint(f)
// 	fmt.Println(x, y, z)
// }

// --------------------------------------------//
// basics/14
// type-inference.go
// package main

// import "fmt"

// func main() {
// 	v := 42.4 + 2i // change me!
// 	fmt.Printf("v is of type %T\n", v)
// }

// --------------------------------------------//
// basics/15
// constants.go
// package main

// import "fmt"

// const Pi = 3.14

// func main() {
// 	const World = "世界"
// 	fmt.Println("Hello", World)
// 	fmt.Println("Happy", Pi, "Day")

// 	const Truth = true
// 	fmt.Println("Go rules?", Truth)
// }

// --------------------------------------------//
// basics/16
// numeric-constants.go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
