package raytracer

import (
	"github.com/EndIsNear/Go-Raytracer/src/mymath"
	"github.com/EndIsNear/Go-Raytracer/src/utils"
	"image"
)

type Texture interface {
	GetColor(*IntersectionData) utils.Color
}

type Shader interface {
	Shade(*IntersectionData, utils.Color, *Scene) utils.Color
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

type TextureImg struct {
	width, height int
	img           *image.Image
}

//for now loads only from png
func InitTexture(filepath string) (bool, *TextureImg) {
	var res TextureImg
	var ok bool
	ok, res.img = utils.OpenPNG(filepath)
	if !ok {
		return false, nil
	}
	res.width = (*res.img).Bounds().Dx()
	res.height = (*res.img).Bounds().Dy()
	return true, &res
}

func (t *TextureImg) GetColor(id *IntersectionData) utils.Color {
	newU := (int(mymath.Floor(id.u)) + t.width*1e3) % t.width
	newV := (int(mymath.Floor(id.v)) + t.height*1e3) % t.height
	res := (*t.img).At(newU, newV)
	r, g, b, _ := res.RGBA()
	return utils.NewColor(uint8(r), uint8(g), uint8(b))
}

type Lambert struct {
	text Texture
}

func (c *Lambert) Shade(id *IntersectionData, ambLight utils.Color, scene *Scene) utils.Color {
	res := utils.ColorMultiplication(c.text.GetColor(id), ambLight)

	for _, light := range scene.lights {
		nearPnt := mymath.VectorsAddition(id.pos, mymath.VectorFloatMultiply(id.normal, 1e-5))
		if TestVisibility(light.pos, nearPnt, scene) {
			lightDir := mymath.VectorsSubstraction(light.pos, id.pos)
			mult := lightDir.LenghtSqr()
			lightDir.Normalize()
			cosTheta := mymath.VectorsDotProduct(lightDir, id.normal)
			lightContr := utils.ColorAddition(ambLight, utils.ColorFloatMulti(*(light.color), light.power/mult*cosTheta))
			res = utils.ColorMultiplication(res, lightContr)
		} else {
			res = utils.ColorMultiplication(res, utils.Color{0.5, 0.5, 0.5})
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
