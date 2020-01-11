package comb

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Nchoosek(n, k int) int {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
