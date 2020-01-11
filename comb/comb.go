package comb

var cache map[int]int

func init() {
	cache = make(map[int]int)
}

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

func Combinations(superset []interface{}, n int) [][]interface{} {
	if len(superset) == n {
		return [][]interface{}{superset}
	}
	if n == 1 {
		res := make([][]interface{}, len(superset))
		for i, m := range res {
			res[i] = []interface{}{m}
		}
		return res
	}
	res := make([][]interface{}, 0)
	for i, m := range superset {
		if i > len(superset)-n {
			break
		}
		others := superset[i+1:]
		combs := Combinations(others, n-1)
		for _, c := range combs {
			set := append([]interface{}{m}, c)
			res = append(res, set)
		}
	}
	return res
}
