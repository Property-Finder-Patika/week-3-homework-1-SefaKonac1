package main

import (
	"fmt"
	"math"
)

func main() {
	num := 35
	calculatePrimeFactors(num)

}

/*Prime factor algorithm was taken by */
func calculatePrimeFactors(n int) {

	for n%2 == 0 {
		fmt.Print("2 ")
		n /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			fmt.Printf(" %d ", i)
			n /= i
		}
	}
}
