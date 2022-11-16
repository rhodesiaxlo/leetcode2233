package main

import "fmt"

func main() {
	fmt.Println(getPrimeNumberCount(2))
}
func getPrimeNumberCount(upper int) int {
	isP := make([]bool, upper+1)

	// sieve
	for i := 2; i*i < upper; i++ {
		if !isP[i] {
			for j := 2 * i; j <= upper; j += i {
				isP[j] = true
			}

		}
	}

	ans := 0
	for i := 2; i < upper; i++ {
		if !isP[i] {
			ans++
		}
	}
	return ans
}
