package math

import "math"

// Fib emits an infinite stream of Fibonacci numbers
func Fib() <-chan int {
	ch := make(chan int)
	go func() {
		x := 0
		y := 1
		ch <- x
		ch <- y
		for {
			x, y = y, x+y
			ch <- y
		}
	}()
	return ch
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}