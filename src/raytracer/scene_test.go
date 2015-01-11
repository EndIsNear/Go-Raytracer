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
	test.AddLight(utils.NewColor(255, 255, 255), 10000)
	test.AddLight(utils.NewColor(100, 100, 100), 100)
	test.AddLight(utils.NewColor(5, 5, 5), 1)

	r, g, b := test.Lights[0].color.ToRGB()
	if r != 255 || g != 255 || b != 255 || test.Lights[0].power != 10000.0 {
		t.Errorf("AddLight failed!")
	}

	r, g, b = test.Lights[1].color.ToRGB()
	if r != 100 || g != 100 || b != 100 || test.Lights[1].power != 100.0 {
		t.Errorf("AddLight failed!")
	}

	r, g, b = test.Lights[2].color.ToRGB()
	if r != 5 || g != 5 || b != 5 || test.Lights[2].power != 1.0 {
		t.Errorf("AddLight failed!")
	}
}

func TestAddGeometry(t *testing.T) {
	test := NewScene()
	pl := Plane{XZ, mymath.Vector{1, 0, 0}, 2}
	test.AddGeometry(&pl, "test")
	pl = Plane{XY, mymath.Vector{0, 2, 0}, 3}
	test.AddGeometry(&pl, "test1")
	pl = Plane{YZ, mymath.Vector{0, 0, 3}, 5}
	test.AddGeometry(&pl, "test2")
	if _, ok := test.Geometries["test"]; !ok {
		t.Errorf("AddGeometry failed!")
	}
	if _, ok := test.Geometries["test1"]; !ok {
		t.Errorf("AddGeometry failed!")
	}
	if _, ok := test.Geometries["test2"]; !ok {
		t.Errorf("AddGeometry failed!")
	}
}

func TestAddShader(t *testing.T) {
	test := NewScene()
	ch := Checker{utils.NewColor(0, 0, 0), utils.NewColor(0, 0, 0), 20}
	test.AddShader(&ch, "test")
	ch = Checker{utils.NewColor(0, 0, 0), utils.NewColor(0, 0, 0), 10}
	test.AddShader(&ch, "test1")
	ch = Checker{utils.NewColor(0, 0, 0), utils.NewColor(0, 0, 0), 30}
	test.AddShader(&ch, "test2")
	if _, ok := test.Shaders["test"]; !ok {
		t.Errorf("AddShader failed!")
	}
	if _, ok := test.Shaders["test1"]; !ok {
		t.Errorf("AddShader failed!")
	}
	if _, ok := test.Shaders["test2"]; !ok {
		t.Errorf("AddShader failed!")
	}
}

func TestAddSceneElement(t *testing.T) {
	test := NewScene()
	pl := Plane{XZ, mymath.Vector{0, 0, 0}, 2}
	test.AddGeometry(&pl, "geometry")
	ch := Checker{utils.NewColor(0, 0, 0), utils.NewColor(0, 0, 0), 20}
	test.AddShader(&ch, "shader")

	if ok := test.AddSceneElement("geometry", "shader"); !ok {
		t.Errorf("AddSceneElement failed!")
	}

	if ok := test.AddSceneElement("geometrY", "shader"); ok {
		t.Errorf("AddSceneElement failed!")
	}

	if ok := test.AddSceneElement("geometry", "shadeR"); ok {
		t.Errorf("AddSceneElement failed!")
	}
}
