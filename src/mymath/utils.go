package mymath

import "math"

const (
	DegToRad = 0.017453292519943295769236907684886127134428718885417 // Pi/180
	RadToDeg = 57.295779513082320876798154814105170332405472466564   // 180/Pi
)

func Floor(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}
