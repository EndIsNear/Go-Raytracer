package main

import (
	"Go-Raytracer/src/sdlwrapper"
	"Go-Raytracer/src/utils"
)

var winWidth, winHeight int = 800, 600

func main() {
	disp, _ := sdlwrapper.NewDisplay(winWidth, winHeight)

	col := utils.Color{255, 0, 0}

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			disp.DrawPixel(i, j, &col)
		}
	}

	disp.Flip()

	sdlwrapper.RunWhileExit()

	disp.Destroy()
}
