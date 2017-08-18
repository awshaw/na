package linearalgebra

import (
  "math"
  "math/rand"

  "github.com/dirtytoeknee/na/src/kahan"
)

// List of floats is essentially a vector
type Vector []float64

// Create a random vector
func RandVec(size int) Vector {
  vec := make(Vector, size)
  for i := range vec {
    // entries "normally distributed with standard normal distribution"
    vec[i] = rand.NormFloat64()
    // vec[i] = rand.Float64()
  }
  return vec
}

// duplicate/copy vector
func (vec Vector) DupVec() Vector {
  dup := make(Vector, len(vec))
  copy(dup, vec)
  return dup
}

// vector scalar multiplication
func (vec Vector) ScalarMult(c float64) Vector {
  for i, x := range vec { // i: index, x: value of v[i]
    vec[i] = x * c
  }
  return vec
}

// vector addition
func (vec Vector) VecAdd(vec_ Vector) Vector {
  for i, x := range vec_ {
    vec[i] += x
  }
  return vec
}

// Dot product of two vectors
func (vec Vector) Dot(vec_ Vector) float64 {
  // ensure dimensions of both vectors are equivalent
  if len(vec) != len(vec_) {
    panic("Dimensions are not equivalent.\n")
  }

  // use Kahan summation for accuracy
  vsum := kahan.NewSum64()
  for i, x := range vec {
    vsum.Add(x * vec_[i])
  }
  return vsum.Sum()
}

// faster dot product
func (vec Vector) FastDot(vec_ Vector) float64 {
  if len(vec) != len(vec_) {
    panic("Dimensions are not equivalent.\n")
  }


  var vsum float64
  for i, x := range vec {
    vsum += x * vec_[i]
  }
  return vsum
}

// vector magnitude
func (vec Vector) Magnitude() float64 {
  return math.Sqrt(vec.Dot(vec))
}

// vector max (absval) component
func (vec Vector) MaxAbsComp() float64 {
  var mac float64
  for _, x := range vec {
    mac = math.Max(mac, math.Abs(x))
  }
  return mac
}

// vector max value
func (vec Vector) Max() (float64, int) {
  if len(vec) == 0 {
    return 0, 0
  }
  max := vec[0]
  index := 0
  for i := 1; i < len(vec); i++ {
    if vec[i] > max {
      max = vec[i]
      index = i
    }
  }
  return max, index
}

// vector ma\in value
func (vec Vector) Min() (float64, int) {
  if len(vec) == 0 {
    return 0, 0
  }
  min := vec[0]
  index := 0
  for i := 1; i < len(vec); i++ {
    if vec[i] < min {
      min = vec[i]
      index = i
    }
  }
  return min, index
}
