package main

import (
	"fmt"
	"math/rand"
	"math"

	"github.com/dirtytoeknee/na/src/kahan"
)

const SeriesSize = 100000

func main() {
	rnums := make([]float64, SeriesSize)
	perm := rand.Perm(len(rnums)) // generate random numbers

	for i := range rnums {
		rnums[perm[i]] = float64(i) / 13
	}

	// fmt.Println(nums)

	n := SumNaive(rnums)
	k := kahan.SumAll(rnums)

	fmt.Printf("Naive sum: %.10f\n", n)
	fmt.Printf("Kahan sum: %.10f\n", k)

	diff := math.Abs((n-k))
	fmt.Printf("Difference: %.10f\n", diff)

}

func SumNaive(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}
