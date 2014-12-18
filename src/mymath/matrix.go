package mymath

import (
	"math"
)

type Matrix [3][3]float64

func NewMatrix(diagonal float64) Matrix {
	var res Matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				res[i][j] = diagonal
			} else {
				res[i][j] = 0
			}
		}
	}
	return res
}

func MultiplyVectMatr(v Vector, m Matrix) Vector {
	return Vector{v.X*m[0][0] + v.Y*m[1][0] + v.Z*m[2][0],
		v.X*m[0][1] + v.Y*m[1][1] + v.Z*m[2][1],
		v.X*m[0][2] + v.Y*m[1][2] + v.Z*m[2][2]}
}

func MatrixMultiplication(left, right Matrix) Matrix {
	res := NewMatrix(0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				res[i][j] += left[i][k] * right[k][j]
			}
		}
	}
	return res
}

func (m *Matrix) Determinant() float64 {
	positive := m[0][0]*m[1][1]*m[2][2] + m[0][1]*m[1][2]*m[2][0] + m[0][2]*m[1][0]*m[2][1]
	negative := m[0][0]*m[1][2]*m[2][1] + m[0][1]*m[1][0]*m[2][2] + m[0][2]*m[1][1]*m[2][0]
	return positive - negative
}

//angle should be in radians
func RotationAroundX(angle float64) Matrix {
	res := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	res[1][1] = cos
	res[2][1] = sin
	res[1][2] = -sin
	res[2][2] = cos

	// fmt.Println(sin, cos, angle)
	// fmt.Printf("%f %f %f %f %f %f %f %f %f", res[0][0], res[0][1], res[0][2], res[1][0], res[1][1], res[1][2], res[2][0], res[2][1], res[2][2])
	return res
}

func RotationAroundY(angle float64) Matrix {
	res := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	res[0][0] = cos
	res[2][0] = -sin
	res[0][2] = sin
	res[2][2] = cos

	return res
}

func RotationAroundZ(angle float64) Matrix {
	res := NewMatrix(1)
	sin := math.Sin(angle)
	cos := math.Cos(angle)

	res[0][0] = cos
	res[1][0] = sin
	res[0][1] = -sin
	res[1][1] = cos

	return res
}
