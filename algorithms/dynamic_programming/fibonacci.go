package dynamicprogramming

func FibonacciDyn(n int) int {
	cache := make(map[int]int)
	return fib(n, cache)
}

func fib(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}
	if n < 2 {
		return 1
	}
	res := fib(n-1, cache) + fib(n-2, cache)
	cache[n] = res
	return res
}
