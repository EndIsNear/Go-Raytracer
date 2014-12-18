package mymath

import (
	"math"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	res := Matrix{{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1}}
	matr := NewMatrix(1)
	if res != matr {
		t.Errorf("NewMatrix failed!")
	}

	res[0][0] = 0
	res[1][1] = 0
	res[2][2] = 0
	matr = NewMatrix(0)
	if res != matr {
		t.Errorf("NewMatrix failed!")
	}
}

func TestMatrixMultiplication(t *testing.T) {
	left := NewMatrix(1)
	right := NewMatrix(1)
	needed := NewMatrix(1)
	res := MatrixMultiplication(left, right)
	if needed != res {
		t.Errorf("MatrixMultiplication failed!")
	}

	left = Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	right = Matrix{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	needed = Matrix{{30, 24, 18}, {84, 69, 54}, {138, 114, 90}}
	res = MatrixMultiplication(left, right)

	if needed != res {
		t.Errorf("MatrixMultiplication failed!")
	}

	left = Matrix{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	right = Matrix{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	needed = Matrix{{14, 14, 14}, {14, 14, 14}, {14, 14, 14}}
	res2 := MatrixMultiplication(left, right)

	if needed != res2 {
		t.Errorf("MatrixMultiplication failed!")
	}
}

func TestMatrVectMultiplication(t *testing.T) {
	vect := NewVector(1, 2, 3)
	matr := NewMatrix(1)
	needed := NewVector(1, 2, 3)
	res := MultiplyVectMatr(vect, matr)
	if needed != res {
		t.Errorf("MultiplyVectMatr failed!")
	}

	vect.Set(3, 2, 3)
	matr = Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	needed.Set(32, 40, 48)
	res = MultiplyVectMatr(vect, matr)
	if needed != res {
		t.Errorf("MultiplyVectMatr failed!")
	}
}

func TestDeterminant(t *testing.T) {
	matr := NewMatrix(1)
	if matr.Determinant() != 1 {
		t.Errorf("Determinant failed!")
	}

	matr = Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if matr.Determinant() != 0 {
		t.Errorf("Determinant failed!")
	}
}

func TestRotateX(t *testing.T) {
	vec := NewVector(0, 1, 0)
	needed := NewVector(0, 0, -1)
	rotMatr := RotationAroundX((90.0 / 180.0) * math.Pi)
	vec = MultiplyVectMatr(vec, rotMatr)
	if math.Abs(vec.X-needed.X) > 1e-10 || math.Abs(vec.Y-needed.Y) > 1e-10 || math.Abs(vec.Z-needed.Z) > 1e-10 {
		t.Errorf("RotationAroundX failed!")
	}
}

func TestRotateY(t *testing.T) {
	vec := NewVector(1, 0, 0)
	needed := NewVector(0, 0, 1)
	rotMatr := RotationAroundY((90.0 / 180.0) * math.Pi)
	vec = MultiplyVectMatr(vec, rotMatr)
	if math.Abs(vec.X-needed.X) > 1e-10 || math.Abs(vec.Y-needed.Y) > 1e-10 || math.Abs(vec.Z-needed.Z) > 1e-10 {
		t.Errorf("RotationAroundY failed!")
	}
}

func TestRotateZ(t *testing.T) {
	vec := NewVector(1, 0, 0)
	needed := NewVector(0, -1, 0)
	rotMatr := RotationAroundZ((90.0 / 180.0) * math.Pi)
	vec = MultiplyVectMatr(vec, rotMatr)
	if math.Abs(vec.X-needed.X) > 1e-10 || math.Abs(vec.Y-needed.Y) > 1e-10 || math.Abs(vec.Z-needed.Z) > 1e-10 {
		t.Errorf("RotationAroundZ failed!")
	}
}
