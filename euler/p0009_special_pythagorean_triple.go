package euler

import (
	"github.com/doozr/geuler/euler/assert"
	"github.com/doozr/geuler/euler/math"
)

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.


Answer:
    31875000
 */

func mn(num int) (int, int) {
	for m := 1; m < num; m++ {
		for n := 1; n < m; n++ {
			if m * (m + n) == int(num / 2) {
				return m, n
			}
		}
	}
	return 0, 0
}

func pythagorasTriangleWithPerimeter(p int) (int, int, int) {
	m, n := mn(p)
	a := math.PowInt(m, 2) - math.PowInt(n, 2)
	b := 2 * m * n
	c := math.PowInt(m, 2) + math.PowInt(n, 2)
	return a, b, c
}

func specialPythagoreanTriple() int {
	a, b, c := pythagorasTriangleWithPerimeter(1000)
	return a * b * c
}

func Problem0009SpecialPythagoreanTriple() (string, int, error) {
	name := "Special Pythagorean triple"
	expected := 31875000
	actual := specialPythagoreanTriple()
	return name, actual, assert.Equal(expected, actual)
}
