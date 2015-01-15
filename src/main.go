package main

import (
	"Go-Raytracer/src/raytracer"
	"fmt"
	"os"
)

const (
	help = `
-s  "filepath" read scene from filepath
-d - display rendering
-o "filepath" save image to filepath`
	tooFewArg                  = `Too few arguments.`
	errParseArg                = `Can't parse args.`
	winWidth, winHeight uint16 = 800, 600
)

func main() {
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
		for render.GetState() == raytracer.RENDERING {
			//display progress
		}
	}

	if saveImage {
		//save the result
		outputFilepath = outputFilepath
	}
}
