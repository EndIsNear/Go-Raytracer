package main

import (
	"Go-Raytracer/src/raytracer"
	"Go-Raytracer/src/sdlwrapper"
	"fmt"
	"os"
	"runtime"
)

const (
	help = `
-s  "filepath" read scene from filepath
-d - display rendering
-o "filepath" save image to filepath`
	tooFewArg                  = `Too few arguments.`
	errParseArg                = `Can't parse args.`
	errCreateDisp              = `Can't initialize window.`
	windowName                 = `GoTracer`
	winWidth, winHeight uint16 = 800, 600
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sceneFilepath := ""
	outputFilepath := ""
	displayRendering := false
	saveImage := false
	argc := len(os.Args)
	if argc <= 1 {
		fmt.Println(tooFewArg)
		fmt.Println(help)
	}

	for i := 1; i < argc; i++ {
		if os.Args[i] == "-d" {
			displayRendering = true
		} else if os.Args[i] == "-o" {
			saveImage = true
			i++
			outputFilepath = os.Args[i]
		} else if os.Args[i] == "-s" {
			i++
			sceneFilepath = os.Args[i]
		} else {
			fmt.Println(errParseArg)
			fmt.Printf("Can't parse %s .\n", os.Args[i])
			fmt.Println(help)
		}
	}

	render := raytracer.NewRenderManager(winWidth, winHeight, 0)
	render.InitScene(sceneFilepath)
	render.StartRendering()

	if displayRendering {
		disp, err := sdlwrapper.NewDisplay(int(winWidth), int(winHeight), windowName)
		if err != nil {
			fmt.Println(errCreateDisp)
			return
		}
		defer disp.Destroy()

		//refreshes display while rendering
		for !sdlwrapper.CheckForExitEvent() && render.GetState() == raytracer.RENDERING {
			RefreshDisplay(render, disp)
		}

		//update window title with render time
		disp.SetTitle(windowName + " [render time:" + render.GetRenderTime().String() + "]")
	}

	if saveImage {
		//save the result
		outputFilepath = outputFilepath
	}

	//loop while user close it
	for !sdlwrapper.CheckForExitEvent() {
	}
}

func RefreshDisplay(render *raytracer.RenderManager, disp *sdlwrapper.Display) {
	for i := uint16(0); i < winWidth; i++ {
		for j := uint16(0); j < winHeight; j++ {
			disp.DrawPixel(int(i), int(j), render.GetPixel(i, j))
		}
	}
	disp.Flip()
}
