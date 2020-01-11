package comb

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

func Ncomb(n, m int) int {
	return Fact(n) / (Fact(m) * Fact(n-m))
}
