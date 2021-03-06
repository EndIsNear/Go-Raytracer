package raytracer

import (
	"github.com/EndIsNear/Go-Raytracer/src/mymath"
	"math"
)

type Camera struct {
	upLeft, upRight, downLeft, start mymath.Vector
}

func NewCamera(start mymath.Vector, yaw, pitch, roll, fov, aspect float64) *Camera {
	var tmp Camera
	tmp.start = start
	var x, y, lenXY, wantedLen, scaling float64
	x = -aspect
	y = 1

	corner := mymath.Vector{x, y, 1}
	center := mymath.Vector{0, 0, 1}

	xy := mymath.VectorsSubstraction(corner, center)
	lenXY = xy.Lenght()
	wantedLen = math.Tan(mymath.DegToRad * (fov / 2))
	scaling = wantedLen / lenXY

	x *= scaling
	y *= scaling

	tmp.upLeft = mymath.Vector{x, y, 1}
	tmp.upRight = mymath.Vector{-x, y, 1}
	tmp.downLeft = mymath.Vector{x, -y, 1}

	rotAroundX := mymath.RotationAroundX(mymath.DegToRad * pitch)
	rotAroundY := mymath.RotationAroundY(mymath.DegToRad * yaw)
	rotAroundZ := mymath.RotationAroundZ(mymath.DegToRad * roll)

	rotation := mymath.MatrixMultiplication((mymath.MatrixMultiplication(rotAroundX, rotAroundY)), rotAroundZ)

	tmp.upLeft = mymath.MultiplyVectMatr(tmp.upLeft, rotation)
	tmp.upRight = mymath.MultiplyVectMatr(tmp.upRight, rotation)
	tmp.downLeft = mymath.MultiplyVectMatr(tmp.downLeft, rotation)

	tmp.upLeft.Add(tmp.start)
	tmp.upRight.Add(tmp.start)
	tmp.downLeft.Add(tmp.start)

	return &tmp
}

func (c *Camera) GetRayAt(x, y, scrWidth, scrHeight uint16) mymath.Ray {
	dir := c.upLeft

	width := mymath.VectorsSubstraction(c.upRight, c.upLeft)
	width.Multiply(float64(x) / float64(scrWidth))
	height := mymath.VectorsSubstraction(c.downLeft, c.upLeft)
	height.Multiply(float64(y) / float64(scrHeight))

	dir.Add(width)
	dir.Add(height)

	dir = mymath.VectorsSubstraction(dir, c.start)
	dir.Normalize()

	return mymath.Ray{c.start, dir}
}
