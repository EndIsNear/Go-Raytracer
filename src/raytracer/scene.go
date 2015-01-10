package raytracer

import (
	"Go-Raytracer/src/utils"
)

type SceneElement struct {
	geometry *Geometry
	shader   *Shader
}

type Light struct {
	color *utils.Color
	power float64
}

type Scene struct {
	Geometries map[string]*Geometry
	Shaders    map[string]*Shader
	Elements   []*SceneElement
	Lights     []Light
	//background color
}

func NewScene() *Scene {
	var tmp Scene
	tmp.Geometries = make(map[string]*Geometry)
	tmp.Shaders = make(map[string]*Shader)
	tmp.Elements = make([]*SceneElement, 0)
	tmp.Lights = make([]Light, 0)
	return &tmp
}

func (s *Scene) AddLight(lightCol utils.Color, lightPower float64) {
	s.Lights = append(s.Lights, Light{&lightCol, lightPower})
}

func (s *Scene) AddGeometry(geometry Geometry, geometryName string) {
	s.Geometries[geometryName] = &geometry
}

func (s *Scene) AddShader(shader Shader, shaderName string) {
	s.Shaders[shaderName] = &shader
}

func (s *Scene) AddSceneElement(geometryName, shaderName string) bool {
	var sceneElem SceneElement
	var ok bool
	sceneElem.geometry, ok = s.Geometries[geometryName]
	if !ok {
		return false
	}
	sceneElem.shader, ok = s.Shaders[shaderName]
	if !ok {
		return false
	}
	s.Elements = append(s.Elements, &sceneElem)

	return true
}
