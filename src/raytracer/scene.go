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
	geometries map[string]*Geometry
	shaders    map[string]*Shader
	elements   []*SceneElement
	lights     []Light
	//background color
}

func NewScene() *Scene {
	var tmp Scene
	tmp.geometries = make(map[string]*Geometry)
	tmp.shaders = make(map[string]*Shader)
	tmp.elements = make([]*SceneElement, 0)
	tmp.lights = make([]Light, 0)
	return &tmp
}

func (s *Scene) AddLight(lightCol utils.Color, lightPower float64) {
	s.lights = append(s.lights, Light{&lightCol, lightPower})
}

func (s *Scene) AddGeometry(geometry Geometry, geometryName string) {
	s.geometries[geometryName] = &geometry
}

func (s *Scene) AddShader(shader Shader, shaderName string) {
	s.shaders[shaderName] = &shader
}

func (s *Scene) AddSceneElement(geometryName, shaderName string) bool {
	var sceneElem SceneElement
	var ok bool
	sceneElem.geometry, ok = s.geometries[geometryName]
	if !ok {
		return false
	}
	sceneElem.shader, ok = s.shaders[shaderName]
	if !ok {
		return false
	}
	s.elements = append(s.elements, &sceneElem)

	return true
}
