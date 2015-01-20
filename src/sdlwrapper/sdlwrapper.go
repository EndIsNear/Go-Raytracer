package sdlwrapper

import (
	"Go-Raytracer/src/utils"
	"errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewDisplay(width, height int, name string) (*Display, error) {
	var err error
	var disp Display
	disp.window, err = sdl.CreateWindow(name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.New("Can not create window.")
	}

	disp.renderer, err = sdl.CreateRenderer(disp.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		disp.window.Destroy()
		return nil, errors.New("Can not create renderer.")
	}
	disp.renderer.Clear()

	return &disp, nil
}

func (disp *Display) DrawPixel(x, y int, col *utils.Color) {
	r, g, b := col.ToRGB()
	disp.renderer.SetDrawColor(r, g, b, 255)
	disp.renderer.DrawPoint(x, y)
}

func (disp *Display) Destroy() {
	disp.renderer.Destroy()
	disp.window.Destroy()
}

func (disp *Display) Flip() {
	disp.renderer.Present()
}

func (disp *Display) SetTitle(newTitle string) {
	disp.window.SetTitle(newTitle)
}

func CheckForExitEvent() bool {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return true
		default:
			_ = t
			return false
		}
	}
	return false
}
