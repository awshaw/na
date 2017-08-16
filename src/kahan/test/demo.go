package main

import (
	"fmt"
	"math/rand"
	"kahan"
)

const SeriesSize = 100000000

func main() {
	nums := make([]float64, SeriesSize)
	perm := rand.Perm(len(nums))
	for i := range nums {
		nums[perm[i]] = float64(i) / 13
	}
	fmt.Printf("Naive sum: %.10f\n", naiveSum(nums))
	fmt.Printf("Kahan sum: %.10f\n", kahan.RollSum64(nums))
}

func naiveSum(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}
