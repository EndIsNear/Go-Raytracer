package raytracer

import (
	"github.com/EndIsNear/Go-Raytracer/src/mymath"
	"github.com/EndIsNear/Go-Raytracer/src/utils"
)

type SceneElement struct {
	geometry *Geometry
	shader   *Shader
}

type Light struct {
	color *utils.Color
	power float64
	pos   mymath.Vector
}

type Scene struct {
	elements []*SceneElement
	lights   []Light
	//background color
}

func NewScene() *Scene {
	var tmp Scene
	tmp.elements = make([]*SceneElement, 0)
	tmp.lights = make([]Light, 0)
	return &tmp
}

func (s *Scene) AddLight(lightCol utils.Color, lightPower float64, lightPos mymath.Vector) {
	s.lights = append(s.lights, Light{&lightCol, lightPower, lightPos})
}

func (s *Scene) AddSceneElement(geometry Geometry, shader Shader) bool {
	sceneElem := SceneElement{&geometry, &shader}
	s.elements = append(s.elements, &sceneElem)

	return true
}
