package main

import (
	"Go-Raytracer/src/raytracer"
)

var winWidth, winHeight uint16 = 800, 600

func main() {
	ok, render := raytracer.NewRenderManager(winWidth, winHeight)
	if !ok {
		return
	}
	render.InitScene()
	render.StartRendering()
	render.Display()
	render.Destroy()
}
