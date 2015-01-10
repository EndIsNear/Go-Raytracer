package raytracer

import (
	"Go-Raytracer/src/mymath"
	"Go-Raytracer/src/utils"
)

type Shader interface {
	GetColor(*IntersectionData, *[]Light) utils.Color
}

type Checker struct {
	first, second utils.Color
	size          float64
}

func (c *Checker) GetColor(id *IntersectionData, lights *[]Light) utils.Color {
	//naive, no lights
	sq := (mymath.Floor(id.u/c.size) + mymath.Floor(id.v/c.size)) % 2
	if sq == 0 {
		return c.first
	} else {
		return c.second
	}
}

type Lambert struct {
}

func (c *Lambert) GetColor(id *IntersectionData) *utils.Color {
	return nil
}

type Phong struct {
}

func (c *Phong) GetColor(id *IntersectionData) *utils.Color {
	return nil
}
