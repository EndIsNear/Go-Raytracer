package mymath

import (
	// "fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

type Ray struct {
	start, dir Vector
}

func NewVector(x, y, z float64) Vector {
	return Vector{x, y, z}
}

func (v *Vector) Set(_x, _y, _z float64) {
	v.X = float64(_x)
	v.Y = float64(_y)
	v.Z = float64(_z)
}

func (v *Vector) Lenght() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Multiply(multiplier float64) {
	v.X *= multiplier
	v.Y *= multiplier
	v.Z *= multiplier
}

func (v *Vector) Add(right Vector) {
	*v = VectorsAddition(*v, right)
}

func (v *Vector) UnaryMinus() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vector) Normalize() {
	multiplier := 1 / v.Lenght()
	v.Multiply(multiplier)
}

func VectorsAddition(left, right Vector) Vector {
	return Vector{left.X + right.X, left.Y + right.Y, left.Z + right.Z}
}

//cross product
func VectorsMultiplication(left, right Vector) Vector {
	return Vector{left.Y*right.Z - left.Z*right.Y, left.Z*right.X - left.X*right.Z, left.X*right.Y - left.Y*right.X}
}

func VectorsDotProduct(left, right Vector) float64 {
	return left.X*right.X + left.Y*right.Y + left.Z*right.Z
}

func VectorsSubstraction(left, right Vector) Vector {
	return Vector{left.X - right.X, left.Y - right.Y, left.Z - right.Z}
}
