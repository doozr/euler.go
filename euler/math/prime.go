package math

func PrimeFactors(n int) []int {
	factors := make([]int, 2)
	if n == 2 || n == 3 {
		factors = append(factors, n)
		return factors
	}
	for i := 2; i*i < n; i++ {
		for n%i == 0 {
			n = n / i
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
