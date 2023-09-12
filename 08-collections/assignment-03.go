/* Write a function that returns all the prime numbers beween the given start and end */

package main

import "fmt"

func main() {
	primes := findPrimes(2, 100)
	fmt.Println(primes)
}

func findPrimes(start, end int) []int {
	result := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			result = append(result, no)
		}
	}
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
