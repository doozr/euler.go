package math

func PrimeFactors(n int) []int {
	factors := make([]int, 2)
	if n == 2 || n == 3 {
		factors = append(factors, n)
		return factors
	}
	for i := 2; i * i < n; i++ {
		for n % i == 0 {
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
	ch := make(chan int, n / 10)
	go func() {
		marked := make([]bool, n)
		ch <- 2
		for x := 3; x < n; x += 2 {
			if marked[x] {
				continue
			}
			ch <- x
			for y := x; y < n; y += x {
				marked[y] = true
			}
		}
		close(ch)
	}()
	return ch
}
