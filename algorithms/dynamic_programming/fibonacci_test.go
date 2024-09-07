package dynamicprogramming

import "testing"

func TestFibonacci(t *testing.T) {
	input := []int{10, 20, 30}
	want := []int{89, 10946, 1346269}
	for i := range input {
		if FibonacciDyn(input[i]) != want[i] {
			t.Errorf("Want: %d; Got: %d", want[i], FibonacciDyn(input[i]))
		}
	}
}
