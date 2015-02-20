package raytracer

import (
	"Go-Raytracer/src/mymath"
	"Go-Raytracer/src/utils"
)

type Texture interface {
	GetColor(*IntersectionData) utils.Color
}

type Shader interface {
	Shade(*IntersectionData, utils.Color, []Light) utils.Color
}

type Color struct {
	color utils.Color
}

func (c *Color) GetColor(id *IntersectionData) utils.Color {
	return c.color
}

type Checker struct {
	first, second utils.Color
	size          float64
}

func (c *Checker) GetColor(id *IntersectionData) utils.Color {
	sq := (mymath.Floor(id.u/c.size) + mymath.Floor(id.v/c.size)) % 2
	if sq == 0 {
		return c.first
	} else {
		return c.second
	}
}

type Lambert struct {
	text Texture
}

func (c *Lambert) Shade(id *IntersectionData, ambLight utils.Color, lights []Light) utils.Color {
	res := c.text.GetColor(id)

	for _, light := range lights {
		ok := true
		if ok { //TestVis
			lightDir := mymath.VectorsSubstraction(light.pos, id.pos)
			mult := lightDir.LenghtSqr()
			lightDir.Normalize()
			cosTheta := mymath.VectorsDotProduct(lightDir, id.normal)
			lightContr := utils.ColorAddition(ambLight, utils.ColorFloatMulti(*lights[0].color, light.power/mult*cosTheta))
			res = utils.ColorMultiplication(res, lightContr)
		}
	}

	return res
}

// type Phong struct {
// 	text Texture
// }

// func (c *Phong) Shade(id *IntersectionData, ambLight utils.Color, lights []Light) utils.Color {
// 	return c.text.GetColor(id)
// }
