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
	ambientLight  utils.Color
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
	rm.ambientLight = utils.Color{0.1, 0.1, 0.1}
	rm.camera = NewCamera(mymath.Vector{0, 150, 0}, 0, -30, 0, 90, float64(rm.width)/float64(rm.height))

	pl := Plane{XZ, mymath.Vector{0, 0, 0}, 5200}
	sp := Sphere{mymath.Vector{-100, 0, 200}, 50}
	cb := Cube{mymath.Vector{100, 50, 200}, 50}

	ch := Checker{utils.NewColor(15, 0, 15), utils.NewColor(0, 1, 1), 10}
	sh := Lambert{&ch}
	ch2 := Checker{utils.NewColor(15, 15, 15), utils.NewColor(0, 1, 1), 10}
	sh2 := Lambert{&ch2}

	rm.scene.AddSceneElement(&pl, &sh)
	rm.scene.AddSceneElement(&sp, &sh)
	rm.scene.AddSceneElement(&cb, &sh2)
	rm.scene.AddLight(utils.Color{1, 1, 1}, 1500000, mymath.Vector{-200, 400, 200})
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
		rm.dispBuffer[x][y] = (*resNode.shader).Shade(&data, rm.ambientLight, rm.scene.lights)
	} else {
		rm.dispBuffer[x][y] = utils.NewColor(0, 0, 0)
	}
}
