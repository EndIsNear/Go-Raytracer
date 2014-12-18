package mymath

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	vec := NewVector(1, 0, 0)

	if vec.X != 1 || vec.Y != 0 || vec.Z != 0 {
		t.Errorf("NewVector failed!")
	}
}

func TestVectorsAddition(t *testing.T) {
	var vec, vec2, resVec, neededVec Vector

	vec.Set(1, 2, 3)
	vec2.Set(3, 2, 1)
	neededVec.Set(4, 4, 4)
	resVec = VectorsAddition(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsAddition failed!")
	}

	vec.Set(1, 0, 0)
	vec2.Set(-1, 0, -1)
	neededVec.Set(0, 0, -1)
	resVec = VectorsAddition(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsAddition failed!")
	}

	vec.Set(0, 2, 2)
	vec2.Set(1, 3, 0)
	neededVec.Set(1, 5, 2)
	resVec = VectorsAddition(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsAddition failed!")
	}
}

func TestVectorsMultiplication(t *testing.T) {
	var vec, vec2, resVec, neededVec Vector

	vec.Set(1, 1, 1)
	vec2.Set(1, 1, 1)
	neededVec.Set(0, 0, 0)
	resVec = VectorsMultiplication(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsMultyplication failed!")
	}

	vec.Set(3, 7, 1)
	vec2.Set(2, 1, 3)
	neededVec.Set(20, -7, -11)
	resVec = VectorsMultiplication(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsMultyplication failed!")
	}

	vec.Set(0, 0, 1)
	vec2.Set(5, 0, 3)
	neededVec.Set(0, 5, 0)
	resVec = VectorsMultiplication(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsMultyplication failed!")
	}

}

func TestVectorsDotProduct(t *testing.T) {
	var vec, vec2 Vector
	var res, needed float64

	vec.Set(0, 0, 0)
	vec2.Set(0, 0, 0)
	needed = 0
	res = VectorsDotProduct(vec, vec2)
	if needed != res {
		t.Errorf("VectorsDotProduct failed!")
	}

	vec.Set(4, 3, 4)
	vec2.Set(1, 2, 1)
	needed = 14
	res = VectorsDotProduct(vec, vec2)
	if needed != res {
		t.Errorf("VectorsDotProduct failed!")
	}

	vec.Set(1, 0, 7)
	vec2.Set(0, 8, 0)
	needed = 0
	res = VectorsDotProduct(vec, vec2)
	if needed != res {
		t.Errorf("VectorsDotProduct failed!")
	}
}

func TestVectorsSubstraction(t *testing.T) {
	var vec, vec2, resVec, neededVec Vector

	vec.Set(1, 1, 1)
	vec2.Set(1, 1, 1)
	neededVec.Set(0, 0, 0)
	resVec = VectorsSubstraction(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsSubstraction failed!")
	}

	vec.Set(1, 0, 1)
	vec2.Set(0, 5, 0)
	neededVec.Set(1, -5, 1)
	resVec = VectorsSubstraction(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsSubstraction failed!")
	}

	vec.Set(3, 2, 1)
	vec2.Set(1, 2, 3)
	neededVec.Set(2, 0, -2)
	resVec = VectorsSubstraction(vec, vec2)
	if neededVec != resVec {
		t.Errorf("VectorsSubstraction failed!")
	}
}

//Vector methods

func TestVectorSet(t *testing.T) {
	var vec Vector

	vec.Set(1, 0, 0)
	if vec.X != 1 || vec.Y != 0 || vec.Z != 0 {
		t.Errorf("Vector.Set() failed!")
	}
}

func TestLenght(t *testing.T) {
	var vec Vector

	vec.Set(0, 0, 0)
	if vec.Lenght() != 0 {
		t.Errorf("Vetor.Lenght() failed!")
	}

	vec.Set(0, 8, 0)
	if vec.Lenght() != 8 {
		t.Errorf("Vetor.Lenght() failed!")
	}

	vec.Set(0, 4, 3)
	if vec.Lenght() != 5 {
		t.Errorf("Vetor.Lenght() failed!")
	}
}

func TestMultiply(t *testing.T) {
	var vec, vec2 Vector

	vec.Set(0, 0, 0)
	vec2 = vec
	vec.Multiply(10000)
	if vec != vec2 {
		t.Errorf("Vetor.Multiply() failed!")
	}

	vec.Set(1, 1, 1)
	vec2.Set(3, 3, 3)
	vec.Multiply(3)
	if vec != vec2 {
		t.Errorf("Vetor.Multiply() failed!")
	}

	vec.Set(1, 2, 3)
	vec2.Set(10, 20, 30)
	vec.Multiply(10)
	if vec != vec2 {
		t.Errorf("Vetor.Multiply() failed!")
	}
}

func TestAdd(t *testing.T) {
	var vec, vec2, vec3 Vector
	vec.Set(0, 0, 0)
	vec2 = vec
	vec3 = vec
	vec.Add(vec2)
	if vec != vec3 {
		t.Errorf("Vetor.Add() failed!")
	}

	vec.Set(1, 0, 1)
	vec2.Set(2, 2, 0)
	vec3.Set(3, 2, 1)
	vec.Add(vec2)
	if vec != vec3 {
		t.Errorf("Vetor.Add() failed!")
	}

	vec.Set(7, 6, 5)
	vec2.Set(1, 2, 3)
	vec3.Set(8, 8, 8)
	vec.Add(vec2)
	if vec != vec3 {
		t.Errorf("Vetor.Add() failed!")
	}
}

func TestUnaryMinus(t *testing.T) {
	var vec, vec2 Vector

	vec.Set(0, 0, 0)
	vec2 = vec
	vec.UnaryMinus()
	if vec != vec2 {
		t.Errorf("Vetor.UnaryMinus() failed!")
	}

	vec.Set(1, 2, 3)
	vec2.Set(-1, -2, -3)
	vec.UnaryMinus()
	if vec != vec2 {
		t.Errorf("Vetor.UnaryMinus() failed!")
	}

	vec.Set(1, 0, 2)
	vec2.Set(-1, 0, -2)
	vec.UnaryMinus()
	if vec != vec2 {
		t.Errorf("Vetor.UnaryMinus() failed!")
	}

}

func TestNormalize(t *testing.T) {
	var vec, vec2 Vector

	vec.Set(37, 0, 0)
	vec2.Set(1, 0, 0)
	vec.Normalize()
	if vec.X-vec2.X > 1e-10 || vec.Y-vec2.Y > 1e-10 || vec.Z-vec2.Z > 1e-10 {
		t.Errorf("Vetor.Normalize() failed!")
	}

	vec.Set(37, 37, 0)
	vec2.Set(1/math.Sqrt(2), 1/math.Sqrt(2), 0)
	vec.Normalize()
	if vec.X-vec2.X > 1e-10 || vec.Y-vec2.Y > 1e-10 || vec.Z-vec2.Z > 1e-10 {
		t.Errorf("Vetor.Normalize() failed!")
	}

	vec.Set(37, 37, 37)
	vec2.Set(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3))
	vec.Normalize()
	if vec.X-vec2.X > 1e-10 || vec.Y-vec2.Y > 1e-10 || vec.Z-vec2.Z > 1e-10 {
		t.Errorf("Vetor.Normalize() failed!")
	}

}
