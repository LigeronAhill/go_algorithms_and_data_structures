package tuples

import "fmt"

func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func tuples() {
	square, cube := powerSeries(3)
	fmt.Println("Square", square, "Cube", cube)
}
