package raytracer

import (
	"Go-Raytracer/src/mymath"
)

type IntersectionData struct {
	pos, normal mymath.Vector
	dist, u, v  float64
}

type Geometry interface {
	Intersect(*mymath.Ray, float64) (bool, *IntersectionData)
}

type Plane struct {
	y float64
}

func (p *Plane) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	//naive no check for other element
	if ray.Dir.Y >= 0.0 {
		return false, nil
	} else {
		var tmp IntersectionData
		yDiff := ray.Dir.Y
		wantYDiff := ray.Start.Y - p.y
		mult := wantYDiff / -yDiff
		tmp.pos = mymath.VectorsAddition(ray.Dir, ray.Start)
		tmp.pos.Multiply(mult)
		tmp.u = tmp.pos.X
		tmp.v = tmp.pos.Z
		return true, &tmp
	}
}

type Sphere struct {
}

func (s *Sphere) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	return false, nil
}

type Cube struct {
}

func (c *Cube) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	return false, nil
}
