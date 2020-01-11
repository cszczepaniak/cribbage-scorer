package comb

var cache map[int]int

func Factorial(n int) int {
	res, ok := cache[n]
	if ok {
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = n * Factorial(n-1)
	}
	cache[n] = res
	return res
}

func Nchoosek(n, k int) int {
	return Factorial(n) / (Factorial(k) * Factorial(n-k))
}
