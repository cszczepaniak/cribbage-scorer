package comb

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Ncomb(n, m int) int {
	return Factorial(n) / (Factorial(m) * Factorial(n-m))
}
