package main

import (
	"fmt"
	"math"
)

func main() {
	var delta = func(a, b, c int) int {
		return (b * b) - (4 * a * c)
	}

	res, _ := baskhara(4, -3, -1, delta)
	d := delta(4, -3, -1)

	fmt.Printf("Resultado: %d\nDelta: %d", res, d)
}

func baskhara(a, b, c int, d func(int, int, int) int) (int, error) {
	return (-b + int(math.Sqrt(float64(d(a, b, c))))) / (2 * a), nil
}
