package utils

import (
	"math"
	"testing"
)

func TestColor(t *testing.T) {
	color := NewColor(0, 122, 255)
	r, g, b := color.ToRGB()
	color[2] += 23
	if r != 0 || g != 122 || b != 255 {
		t.Errorf("NewColor or ToRGB failed!")
	}
}

func TestColorsAdd(t *testing.T) {
	color := NewColor(77, 7, 12)
	color2 := NewColor(13, 11, 32)
	needed := NewColor(90, 18, 44)
	res := ColorAddition(color, color2)
	if math.Abs(res[0]-needed[0]) > 1e-10 || math.Abs(res[1]-needed[1]) > 1e-10 || math.Abs(res[2]-needed[2]) > 1e-10 {
		t.Errorf("ColorAddition failed!")
	}
}

func TestColorsMulti(t *testing.T) {
	color := NewColor(3, 3, 3)
	color2 := NewColor(255, 255, 255)
	needed := NewColor(3, 3, 3)
	res := ColorMultiplication(color, color2)
	if math.Abs(res[0]-needed[0]) > 1e-10 || math.Abs(res[1]-needed[1]) > 1e-10 || math.Abs(res[2]-needed[2]) > 1e-10 {
		t.Errorf("ColorMultiplication failed!")
	}
}

func TestColorFloatMulti(t *testing.T) {
	color := NewColor(10, 100, 200)
	color.Multiply(2)
	r, g, b := color.ToRGB()
	if r != 20 || g != 200 || b != 255 {
		t.Errorf("ColorFloatMulti failed!")
	}
}

func TestColorFloatDevide(t *testing.T) {
	color := NewColor(10, 100, 200)
	color.Devide(2)
	r, g, b := color.ToRGB()
	if r != 5 || g != 50 || b != 100 {
		t.Errorf("ColorFloatDevide failed!")
	}
}
