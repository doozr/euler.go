package math

// Fib emits an infinite stream of Fibonacci numbers to the channel
func Fib(ch chan int) {
	x := 0
	y := 1
	ch <- x
	ch <- y
	for {
		x, y = y, x+y
		ch <- y
	}
}
