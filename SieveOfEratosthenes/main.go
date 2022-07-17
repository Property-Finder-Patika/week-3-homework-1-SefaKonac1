package main

import "fmt"

func main() {
	num := 100
	sieveOfEratosthenes(num)
}

func sieveOfEratosthenes(n int) {

	/*It keeps prime numbers till n*/
	/*declare prime number with n+1 size*/
	prime := make([]bool, n+1)

	/*initialize prime array to true value*/
	for i := 0; i < n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// Update all multiples of p greater than or
		// equal to the square of it numbers which
		// are multiple of p and are less than p^2
		// are already been marked.

		if prime[p] {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}

	}

	for i := 2; i <= n; i++ {
		if prime[i] {
			fmt.Printf("%d ", i)
		}
	}

}
