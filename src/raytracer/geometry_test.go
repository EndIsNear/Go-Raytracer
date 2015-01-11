package raytracer

import (
	"Go-Raytracer/src/mymath"
	"testing"
)

func TestIntersectionWithPlane(t *testing.T) {
	ray := mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, -1, 0}}
	plane := Plane{XZ, mymath.Vector{0, 0, 0}, 1}

	//just interesection
	if ok, data := plane.Intersect(&ray, 2); !ok || data.u != 0 || data.v != 0 {
		t.Errorf("Intersection with plane failed!")
	}

	//oposite direction
	ray = mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, 1, 0}}
	if ok, _ := plane.Intersect(&ray, 1e99); ok {
		t.Errorf("Intersection with plane failed!")
	}

	//out of plane
	ray = mymath.Ray{mymath.Vector{2, 2, 2}, mymath.Vector{0, -1, 0}}
	if ok, _ := plane.Intersect(&ray, 1e99); ok {
		t.Errorf("Intersection with plane failed!")
	}
}

func TestIntersectionWithSphere(t *testing.T) {

}

func TestIntersectionWithCube(t *testing.T) {

}
