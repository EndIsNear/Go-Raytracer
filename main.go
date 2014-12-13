package main

import (
	"Go-Raytracer/sdlwrapper"
)

var winWidth, winHeight int = 800, 600

func main() {
	disp, _ := sdlwrapper.NewDisplay(winWidth, winHeight)

	disp.SetDrawColor(255, 0, 0, 0)

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			disp.DrawPoint(i, j)
		}
	}

	disp.SetDrawColor(0, 255, 0, 0)

	for i := 50; i < 100; i++ {
		for j := 0; j < 50; j++ {
			disp.DrawPoint(i, j)
		}
	}

	disp.SetDrawColor(0, 0, 255, 0)

	for i := 100; i < 150; i++ {
		for j := 0; j < 50; j++ {
			disp.DrawPoint(i, j)
		}
	}

	disp.Flip()

	sdlwrapper.Sleep(5000)

	disp.Destroy()
}
