package raytracer

import (
	"Go-Raytracer/src/mymath"
	"math"
)

const (
	XY = iota
	XZ
	YZ
)

type IntersectionData struct {
	pos, normal mymath.Vector
	dist, u, v  float64
}

type Geometry interface {
	Intersect(*mymath.Ray, float64) (bool, *IntersectionData)
}

type Plane struct {
	orientation uint8
	center      mymath.Vector
	limit       float64
}

func (p *Plane) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	var (
		start, direction, plane float64
		result                  IntersectionData
	)
	if p.orientation == XY {
		start = ray.Start.Z
		direction = ray.Dir.Z
		plane = p.center.Z
	} else if p.orientation == XZ {
		start = ray.Start.Y
		direction = ray.Dir.Y
		plane = p.center.Y
	} else {
		start = ray.Start.X
		direction = ray.Dir.X
		plane = p.center.X
	}

	if direction >= 0.0 && start > plane || direction <= 0.0 && start < plane {
		return false, nil
	} else {
		wantDiff := start - plane
		mult := wantDiff / -direction

		if crnDist < mult {
			return false, nil
		}

		result.pos = ray.Dir
		result.pos.Multiply(mult)
		result.pos = mymath.VectorsAddition(result.pos, ray.Start)
		result.dist = mult

		if p.orientation == XY {
			if math.Abs(p.center.X-result.pos.X) > p.limit/2 || math.Abs(p.center.Y-result.pos.Y) > p.limit/2 {
				return false, nil
			}
			result.u = result.pos.X
			result.v = result.pos.Y
			result.normal = mymath.Vector{0, 0, 1}
		} else if p.orientation == XZ {
			if math.Abs(p.center.X-result.pos.X) > p.limit/2 || math.Abs(p.center.Z-result.pos.Z) > p.limit/2 {
				return false, nil
			}
			result.u = result.pos.X
			result.v = result.pos.Z
			result.normal = mymath.Vector{0, 1, 0}
		} else {
			if math.Abs(p.center.Y-result.pos.Y) > p.limit/2 || math.Abs(p.center.Z-result.pos.Z) > p.limit/2 {
				return false, nil
			}
			result.u = result.pos.Y
			result.v = result.pos.Z
			result.normal = mymath.Vector{1, 0, 0}
		}

		return true, &result
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
