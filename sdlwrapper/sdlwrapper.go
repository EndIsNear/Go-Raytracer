package sdlwrapper

import (
	"errors"
	"github.com/veandco/go-sdl2/sdl"
)

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewDisplay(width, height int) (*Display, error) {
	var err error
	var disp Display
	disp.window, err = sdl.CreateWindow("My GO tracer", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
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

func (disp Display) SetDrawColor(r, g, b, a uint8) {
	disp.renderer.SetDrawColor(r, g, b, a)
}

func (disp Display) DrawPoint(x, y int) {
	disp.renderer.DrawPoint(x, y)
}

func (disp Display) Destroy() {
	disp.renderer.Destroy()
	disp.window.Destroy()
}

func (disp Display) Flip() {
	disp.renderer.Present()
}

func Sleep(ms uint32) {
	sdl.Delay(ms)
}
