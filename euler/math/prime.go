package math

import (
	"math"
)

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

func Primes(n int) <-chan int {
	ch := make(chan int)
	go func() {
		marked := make([]bool, n)
		sqrtn := int(math.Sqrt(float64(n)))
		for x := 3; x <= sqrtn; x += 2 {
			if !marked[x] {
				for y := x * x; y < n; y += x {
					marked[y] = true
				}
			}
		}
		for x := 1; x < n; x += 2 {
			if !marked[x] {
				ch <- x
			}
		}
		close(ch)
	}()
	return ch
}
