package raytracer

import (
	"Go-Raytracer/src/mymath"
	"Go-Raytracer/src/sdlwrapper"
	"Go-Raytracer/src/utils"
)

type RenderManager struct {
	display       *sdlwrapper.Display
	dispBuffer    [][]utils.Color
	scene         *Scene
	camera        *Camera
	width, height uint16
}

func NewRenderManager(width, height uint16) (bool, *RenderManager) {
	var err error
	var res RenderManager
	res.width = width
	res.height = height

	res.display, err = sdlwrapper.NewDisplay(int(width), int(height))
	if err != nil {
		return false, nil
	}

	res.scene = NewScene()
	res.dispBuffer = make([][]utils.Color, width)
	for i := uint16(0); i < width; i++ {
		res.dispBuffer[i] = make([]utils.Color, height)
	}

	return true, &res
}

func (rm *RenderManager) InitScene() {
	//hardcoded for now
	rm.camera = NewCamera(mymath.Vector{0, 150, 0}, 0, -10, 0, 90, float64(rm.width)/float64(rm.height))

	pl := Plane{XZ, mymath.Vector{0, 0, 400}, 200}
	rm.scene.AddGeometry(&pl, "plane")
	pl2 := Plane{YZ, mymath.Vector{100, 100, 400}, 200}
	rm.scene.AddGeometry(&pl2, "plane2")
	pl3 := Plane{XY, mymath.Vector{0, 100, 500}, 200}
	rm.scene.AddGeometry(&pl3, "plane3")

	ch := Checker{utils.NewColor(0, 122, 122), utils.NewColor(0, 33, 33), 10}
	rm.scene.AddShader(&ch, "checker")
	ch2 := Checker{utils.NewColor(122, 122, 0), utils.NewColor(33, 33, 0), 20}
	rm.scene.AddShader(&ch2, "checker2")
	ch3 := Checker{utils.NewColor(122, 0, 122), utils.NewColor(33, 0, 33), 30}
	rm.scene.AddShader(&ch3, "checker3")

	rm.scene.AddSceneElement("plane", "checker")
	rm.scene.AddSceneElement("plane2", "checker2")
	rm.scene.AddSceneElement("plane3", "checker3")
}

func (rm *RenderManager) StartRendering() {
	for i := uint16(0); i < rm.width; i++ {
		for j := uint16(0); j < rm.height; j++ {
			rm.raytrace(i, j)
		}
	}
}

func (rm *RenderManager) raytrace(x, y uint16) {
	var (
		ray     mymath.Ray
		data    IntersectionData
		tmp     *IntersectionData
		ok      bool
		resNode *SceneElement
	)

	data.dist = 1e99
	ray = rm.camera.GetRayAt(x, y, rm.width, rm.height)

	for _, val := range rm.scene.Elements {
		ok, tmp = (*val.geometry).Intersect(&ray, data.dist)
		if ok {
			data = *tmp
			resNode = val
		}
	}

	if data.dist < 1e99 {
		rm.dispBuffer[x][y] = (*resNode.shader).GetColor(&data, &rm.scene.Lights)
	} else {
		rm.dispBuffer[x][y] = utils.NewColor(0, 0, 0)
	}
}

func (rm *RenderManager) Display() {
	for i := 0; i < int(rm.width); i++ {
		for j := 0; j < int(rm.height); j++ {
			rm.display.DrawPixel(i, j, &rm.dispBuffer[i][j])
		}
	}
	rm.display.Flip()
	sdlwrapper.RunWhileExit()
}

func (rm *RenderManager) Destroy() {
	rm.display.Destroy()
}
