package raytracer

import (
	"Go-Raytracer/src/mymath"
	"Go-Raytracer/src/utils"
	"time"
)

const (
	WindowName = "GO raytracer."
)

//render states
const (
	NOT_INITIALIZED = iota
	RENDERING
	STOPED
	FINISHED
)

type RenderManager struct {
	dispBuffer    [][]utils.Color
	scene         *Scene
	camera        *Camera
	width, height uint16
	workersCnt    uint16
	startTime     time.Time
	state         uint16
}

func NewRenderManager(width, height, workersCnt uint16) *RenderManager {
	var res RenderManager
	res.width = width
	res.height = height
	res.workersCnt = workersCnt
	res.state = NOT_INITIALIZED

	res.scene = NewScene()
	res.dispBuffer = make([][]utils.Color, width)
	for i := uint16(0); i < width; i++ {
		res.dispBuffer[i] = make([]utils.Color, height)
	}

	return &res
}

func (rm *RenderManager) InitScene(sceneFileName string) {
	//hardcoded for now
	//should be read from file
	rm.camera = NewCamera(mymath.Vector{0, 150, 0}, 0, -10, 0, 90, float64(rm.width)/float64(rm.height))

	pl := Plane{XZ, mymath.Vector{0, 0, 400}, 200}
	pl2 := Plane{YZ, mymath.Vector{100, 100, 400}, 200}
	pl3 := Plane{XY, mymath.Vector{0, 100, 500}, 200}

	ch := Checker{utils.NewColor(0, 122, 122), utils.NewColor(0, 33, 33), 10}
	ch2 := Checker{utils.NewColor(122, 122, 0), utils.NewColor(33, 33, 0), 20}
	ch3 := Checker{utils.NewColor(122, 0, 122), utils.NewColor(33, 0, 33), 30}

	rm.scene.AddSceneElement(&pl, &ch)
	rm.scene.AddSceneElement(&pl2, &ch2)
	rm.scene.AddSceneElement(&pl3, &ch3)
	rm.state = STOPED
}

func (rm *RenderManager) GetPixel(x, y uint16) *utils.Color {
	return &rm.dispBuffer[x][y]
}

func (rm *RenderManager) GetRenderTime() time.Duration {
	return time.Now().Sub(rm.startTime)
}

func (rm *RenderManager) GetWidth() uint16 {
	return rm.height
}

func (rm *RenderManager) GetHeight() uint16 {
	return rm.height
}

func (rm *RenderManager) StartRendering() {
	rm.state = RENDERING
	go rm.rendering()
}

func (rm *RenderManager) StopRendering() {
	rm.state = STOPED
}

func (rm *RenderManager) GetState() uint16 {
	return rm.state
}

func (rm *RenderManager) rendering() {
	rm.startTime = time.Now()
	for i := uint16(0); i < rm.width; i++ {
		for j := uint16(0); j < rm.height; j++ {
			rm.raytrace(i, j)
		}
	}
	rm.state = FINISHED
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

	for _, val := range rm.scene.elements {
		ok, tmp = (*val.geometry).Intersect(&ray, data.dist)
		if ok {
			data = *tmp
			resNode = val
		}
	}

	if data.dist < 1e99 {
		rm.dispBuffer[x][y] = (*resNode.shader).GetColor(&data, &rm.scene.lights)
	} else {
		rm.dispBuffer[x][y] = utils.NewColor(0, 0, 0)
	}
}
