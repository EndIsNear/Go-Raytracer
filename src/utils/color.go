package utils

type color [3]float64

func NewColor(r, g, b uint8) color {
	return color{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

func (c color) ToRGB() (uint8, uint8, uint8) {
	c[0] = between01(c[0])
	c[1] = between01(c[1])
	c[2] = between01(c[2])
	return uint8(c[0] * 255), uint8(c[1] * 255), uint8(c[2] * 255)
}

func (c *color) Multiply(mult float64) {
	*c = ColorFloatMulti(*c, mult)
}

func (c *color) Devide(divider float64) {
	*c = ColorFloatDevide(*c, divider)
}

func ColorAddition(left, right color) color {
	return color{left[0] + right[0], left[1] + right[1], left[2] + right[2]}
}

func ColorMultiplication(left, right color) color {
	return color{left[0] * right[0], left[1] * right[1], left[2] * right[2]}
}

func ColorFloatMulti(c color, mult float64) color {
	return color{c[0] * mult, c[1] * mult, c[2] * mult}
}

func ColorFloatDevide(c color, divider float64) color {
	return color{c[0] / divider, c[1] / divider, c[2] / divider}
}

func between01(color float64) float64 {
	if color < 0 {
		color = 0
	}
	if color > 1 {
		color = 1
	}
	return color
}
