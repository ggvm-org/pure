package pure

func Fib(n int) int {
	if n == 0 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
