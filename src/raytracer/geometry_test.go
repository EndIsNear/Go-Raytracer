package raytracer

import (
	"github.com/EndIsNear/Go-Raytracer/src/mymath"
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
	ray := mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, -1, 0}}
	sphere := Sphere{mymath.Vector{0, 0, 0}, 1}

	//just interesection
	if ok, data := sphere.Intersect(&ray, 2); !ok || (data.pos != mymath.Vector{0, 1, 0}) {
		t.Errorf("Intersection with sphere failed!")
	}

	//oposite dir
	ray = mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, 1, 0}}
	if ok, _ := sphere.Intersect(&ray, 2); ok {
		t.Errorf("Intersection with sphere failed!")
	}

	//inside
	ray = mymath.Ray{mymath.Vector{0, 0, 0}, mymath.Vector{0, 1, 0}}
	if ok, data := sphere.Intersect(&ray, 2); !ok || (data.pos != mymath.Vector{0, 1, 0}) {
		t.Errorf("Intersection with sphere failed!")
	}

	//no intersetion
	ray = mymath.Ray{mymath.Vector{0, 2.1, 2.1}, mymath.Vector{0, -1, 0}}
	if ok, _ := sphere.Intersect(&ray, 2); ok {
		t.Errorf("Intersection with sphere failed!")
	}
}

func TestIntersectionWithCube(t *testing.T) {
	ray := mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, -1, 0}}
	cube := Cube{mymath.Vector{0, 0, 0}, 2}

	//just interesection
	if ok, data := cube.Intersect(&ray, 2); !ok || (data.pos != mymath.Vector{0, 1, 0}) {
		t.Errorf("Intersection with cube failed!")
	}

	//oposite dir
	ray = mymath.Ray{mymath.Vector{0, 2, 0}, mymath.Vector{0, 1, 0}}
	if ok, _ := cube.Intersect(&ray, 2); ok {
		t.Errorf("Intersection with cube failed!")
	}

	//inside
	ray = mymath.Ray{mymath.Vector{0, 0, 0}, mymath.Vector{0, 1, 0}}
	if ok, data := cube.Intersect(&ray, 2); !ok || (data.pos != mymath.Vector{0, 1, 0}) {
		t.Errorf("Intersection with cube failed!")
	}

	//no intersetion
	ray = mymath.Ray{mymath.Vector{0, 2.1, 2.1}, mymath.Vector{0, -1, 0}}
	if ok, _ := cube.Intersect(&ray, 2); ok {
		t.Errorf("Intersection with cube failed!")
	}

}
