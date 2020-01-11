//Combinatorics utils
package comb

//Factorial
func Fact(n int) int {
	if n > 0 {
		return n * Fact(n - 1)
	}
	return 1
}

//Number of combinations - nCr function
func Ncomb(n, m int) int {
	return Fact(n) / (Fact(m) * Fact(n - m))
}