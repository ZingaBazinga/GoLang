package exercise3

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

func TempConvertor() {
	c := 25.0
	f := celsiusToFahrenheit(c)
	fmt.Printf("%.1f°C = %.1f°F\n", c, f)

	f2 := 77.0
	c2 := fahrenheitToCelsius(f2)
	fmt.Printf("%.1f°F = %.1f°C\n", f2, c2)
}
