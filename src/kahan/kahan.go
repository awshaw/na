// Kahan algorithm implemented in Go.
// Summarized from wiki: the Kahan summation algo (aka compensated
// summation) significantly reduces numerical error in the total
// summation of a finite sequence of floats. This is done by keeping
// some sort of variable which accumulates small errors.
// @ github.com/dirtytoeknee

package kahan

// RollSum64 computes the rolling sum of our floats (float64's)
type RollSum64 struct {
  sum float64
  comp float64 // our running compensation
}

// Create a RollSum64 initialized to 0
func NewSum64() *RollSum64 {
  return &RollSum64{}
}

// Add a new number and return the new sum
func (s *RollSum64) Add(n float64) float64 {
  n -= s.comp
  sum := s.sum + n
  s.comp = (sum - s.sum) - n
  s.sum = sum
  return s.sum
}

// Return current sum
func (s *RollSum64) Sum() float64 {
  return s.sum
}

func SumAll(nums []float64) float64 {
  var roll RollSum64
  for _, n := range nums {
    roll.Add(n)
  }
  return roll.Sum()
}
