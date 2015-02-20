package raytracer

import (
	"Go-Raytracer/src/mymath"
	"Go-Raytracer/src/utils"
	"testing"
)

func TestNewScene(t *testing.T) {

	if test := NewScene(); test == nil {
		t.Errorf("NewScene failed!")
	}
}

func TestAddLight(t *testing.T) {
	test := NewScene()
	test.AddLight(utils.NewColor(255, 255, 255), 10000, mymath.Vector{0, 0, 0})
	test.AddLight(utils.NewColor(100, 100, 100), 100, mymath.Vector{0, 0, 0})
	test.AddLight(utils.NewColor(5, 5, 5), 1, mymath.Vector{0, 0, 0})

	r, g, b := test.lights[0].color.ToRGB()
	if r != 255 || g != 255 || b != 255 || test.lights[0].power != 10000.0 {
		t.Errorf("AddLight failed!")
	}

	r, g, b = test.lights[1].color.ToRGB()
	if r != 100 || g != 100 || b != 100 || test.lights[1].power != 100.0 {
		t.Errorf("AddLight failed!")
	}

	r, g, b = test.lights[2].color.ToRGB()
	if r != 5 || g != 5 || b != 5 || test.lights[2].power != 1.0 {
		t.Errorf("AddLight failed!")
	}
}

func TestAddSceneElement(t *testing.T) {
	test := NewScene()
	pl := Plane{XZ, mymath.Vector{0, 0, 0}, 2}
	ch := Checker{utils.NewColor(0, 0, 0), utils.NewColor(0, 0, 0), 20}
	lm := Lambert{&ch}

	if ok := test.AddSceneElement(&pl, &lm); !ok {
		t.Errorf("AddSceneElement failed!")
	}
}
