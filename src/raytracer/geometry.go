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

func TestVisibility(start, end mymath.Vector, scene *Scene) bool {
	rayDir := mymath.VectorsSubstraction(end, start)
	crnDist := rayDir.Lenght()
	rayDir.Normalize()
	ray := mymath.Ray{start, rayDir}

	var ok bool
	for _, val := range scene.elements {
		if ok, _ = (*val.geometry).Intersect(&ray, crnDist); ok {
			return false
		}
	}
	return true
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

		//adjust the normal
		if mymath.VectorsDotProduct(ray.Dir, result.normal) > 0 {
			result.normal.UnaryMinus()
		}
		return true, &result
	}
}

type Sphere struct {
	center mymath.Vector
	radius float64
}

func (s *Sphere) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	H := mymath.VectorsSubstraction(ray.Start, s.center)
	A := ray.Dir.LenghtSqr()
	B := 2 * mymath.VectorsDotProduct(H, ray.Dir)
	C := H.LenghtSqr() - s.radius*s.radius
	dscr := B*B - 4*A*C
	if dscr < 0 {
		return false, nil
	}
	x1 := (-B + math.Sqrt(dscr)) / (2 * A)
	x2 := (-B - math.Sqrt(dscr)) / (2 * A)
	if x1 < 0 && x2 < 0 {
		return false, nil
	} else {
		var res IntersectionData
		res.dist = x2
		if res.dist < 0 {
			res.dist = x1
		}
		res.pos = mymath.VectorsAddition(ray.Start, mymath.VectorFloatMultiply(ray.Dir, res.dist))
		res.normal = mymath.VectorsSubstraction(res.pos, s.center)
		res.normal.Normalize()
		res.u = 0 //(math.Pi + math.Atan2(res.pos.Z-s.center.Z, res.pos.X-s.center.X)) / (2 * math.Pi)
		res.v = 0 //1.0 - (math.Pi/2+math.Asin((res.pos.Y-s.center.Y)/s.radius))/math.Pi

		return true, &res
	}

}

type Cube struct {
	center   mymath.Vector
	sideSize float64
}

func (c *Cube) Intersect(ray *mymath.Ray, crnDist float64) (bool, *IntersectionData) {
	planeCenter := mymath.Vector{0, 0, c.sideSize / 2}
	cubeSide := Plane{XY, mymath.VectorsAddition(c.center, planeCenter), c.sideSize}
	intersect, res := cubeSide.Intersect(ray, crnDist)
	if intersect && crnDist > res.dist {
		crnDist = res.dist
	}

	planeCenter.UnaryMinus()
	cubeSide.center = mymath.VectorsAddition(c.center, planeCenter)
	intersectTmp, tmp := cubeSide.Intersect(ray, crnDist)
	if intersectTmp && crnDist > tmp.dist {
		crnDist = tmp.dist
		intersect, res = intersectTmp, tmp
	}

	planeCenter = mymath.Vector{0, c.sideSize / 2, 0}
	cubeSide = Plane{XZ, mymath.VectorsAddition(c.center, planeCenter), c.sideSize}
	intersectTmp, tmp = cubeSide.Intersect(ray, crnDist)
	if intersectTmp && crnDist > tmp.dist {
		crnDist = tmp.dist
		res = tmp
		intersect = intersectTmp
	}

	planeCenter.UnaryMinus()
	cubeSide.center = mymath.VectorsAddition(c.center, planeCenter)
	intersectTmp, tmp = cubeSide.Intersect(ray, crnDist)
	if intersectTmp && crnDist > tmp.dist {
		crnDist = tmp.dist
		res = tmp
		intersect = intersectTmp
	}

	planeCenter = mymath.Vector{c.sideSize / 2, 0, 0}
	cubeSide = Plane{YZ, mymath.VectorsAddition(c.center, planeCenter), c.sideSize}
	intersectTmp, tmp = cubeSide.Intersect(ray, crnDist)
	if intersectTmp && crnDist > tmp.dist {
		crnDist = tmp.dist
		res = tmp
		intersect = intersectTmp
	}

	planeCenter.UnaryMinus()
	cubeSide.center = mymath.VectorsAddition(c.center, planeCenter)
	intersectTmp, tmp = cubeSide.Intersect(ray, crnDist)
	if intersectTmp && crnDist > tmp.dist {
		crnDist = tmp.dist
		res = tmp
		intersect = intersectTmp
	}

	return intersect, res
}
