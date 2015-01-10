package utils

type Color [3]float64

func NewColor(r, g, b uint8) Color {
	return Color{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

func (c Color) ToRGB() (uint8, uint8, uint8) {
	c[0] = between01(c[0])
	c[1] = between01(c[1])
	c[2] = between01(c[2])
	return uint8(c[0] * 255), uint8(c[1] * 255), uint8(c[2] * 255)
}

func (c *Color) Multiply(mult float64) {
	*c = ColorFloatMulti(*c, mult)
}

func (c *Color) Devide(divider float64) {
	*c = ColorFloatDevide(*c, divider)
}

func ColorAddition(left, right Color) Color {
	return Color{left[0] + right[0], left[1] + right[1], left[2] + right[2]}
}

func ColorMultiplication(left, right Color) Color {
	return Color{left[0] * right[0], left[1] * right[1], left[2] * right[2]}
}

func ColorFloatMulti(c Color, mult float64) Color {
	return Color{c[0] * mult, c[1] * mult, c[2] * mult}
}

func ColorFloatDevide(c Color, divider float64) Color {
	return Color{c[0] / divider, c[1] / divider, c[2] / divider}
}

func between01(Color float64) float64 {
	if Color < 0 {
		Color = 0
	}
	if Color > 1 {
		Color = 1
	}
	return Color
}
