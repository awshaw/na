package linearalgebra

import (
  // "bytes"
  // "strconv"

  "github.com/dirtytoeknee/na/src/kahan"
)

type Matrix struct {
  Rows int
  Cols int
  Data []float64
}

func CreateMatrix(rows, cols int) *Matrix {
  mat := &Matrix{
    Rows: rows,
    Cols: cols,
    Data: make([]float64, rows*cols),
  }
  return mat
}

// Return the element from the ith row & jth col
func (m *Matrix) Elem(i, j int) float64 {
	return m.Data[i*m.Cols+j]
}

// Updates element in ith row & jth col
func (mat *Matrix) Upd(i,j int, updateVal float64) {
    mat.Data[i*mat.Cols + j] = updateVal
}

// Return identity matrix
func IMat(size int) *Matrix {
  mat := CreateMatrix(size, size)
  for i := 0; i < size; i++ {
    mat.Upd(i,i,1)
  }
  return mat
}

// Copy matrix
func (mat *Matrix) CopyMat() *Matrix {
  mat_c := mat
  return mat_c
}

// Return an n x 1 matrix, given a vector
func ColMat(vec Vector) *Matrix {
  mat := CreateMatrix(len(vec), 1)
  copy(mat.Data, vec)
  return mat
}

// Add two matrices of equal dimensions
func (mat *Matrix) Add(mat_ *Matrix) *Matrix {
  if mat.Rows != mat_.Rows || mat.Cols != mat_.Cols {
    panic("Dimensions not equivalent.\n")
  }
  for i, j := range mat_.Data {
    mat.Data[i] += j
  }
  return mat
}

// Multiply two matrices
func (mat *Matrix) Mult(mat_ *Matrix) *Matrix {
  if mat.Cols != mat_.Rows {
    panic("Dimensions not equivalent.\n")
  }
  tmp := &Matrix {
    Rows: mat.Rows,
    Cols: mat_.Cols,
    Data: make([]float64, mat.Rows * mat_.Cols),
  }
  Index := 0
  for i := 0; i < tmp.Rows; i++ {
    for j := 0; j < tmp.Cols; j++ {
      sum := kahan.NewSum64()
      for k := 0; k < mat.Cols; k++ {
        sum.Add(mat.Elem(i, k) * mat_.Elem(k, j))
      }
      tmp.Data[Index] = sum.Sum()
      Index++
    }
  }
  return tmp
}

// Don't use a numerically accurate sum such as the Kahan algo,
// thus increasing performance speeds, but decreasing num. accuracy
func (mat *Matrix) FastMul(mat_ *Matrix) *Matrix {
  if mat.Cols != mat_.Rows {
    panic("Dimensions not equivalent.\n")
  }
  tmp := &Matrix {
    Rows: mat.Rows,
    Cols: mat_.Cols,
    Data: make([]float64, mat.Rows * mat_.Cols),
  }
  Index := 0
  for i := 0; i < tmp.Rows; i++ {
    for j := 0; j < tmp.Cols; j++ {
      var sum float64
      for k := 0; k < mat.Cols; k++ {
        sum += mat.Elem(i, k) * mat_.Elem(k, j)
      }
      tmp.Data[Index] = sum
      Index++
    }
  }
  return tmp
}

// Return transpose of a given matrix
func (mat *Matrix) Transpose() *Matrix {
  tmp := CreateMatrix(mat.Cols, mat.Rows)
  for i := 0; i < mat.Rows; i++ {
    for j:= 0; j < mat.Cols; j++ {
      tmp.Upd(j, i, mat.Elem(i, j))
    }
  }
  return tmp
}
