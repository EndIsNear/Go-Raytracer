package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type PNGSaver struct {
	imgWidth, imgHeight int
	fileName            string
	rect                image.Rectangle
	image               *image.RGBA
}

func NewPNG(imgWidth, imgHeight int, fileName string) *PNGSaver {
	var res PNGSaver
	res.imgHeight = imgHeight
	res.imgWidth = imgWidth
	res.rect = image.Rect(0, 0, imgWidth, imgHeight)
	res.image = image.NewRGBA(res.rect)
	res.fileName = fileName
	return &res
}

func (i *PNGSaver) SetPixelAt(x, y int, col *Color) {
	r, g, b := col.ToRGB()
	i.image.Set(x, y, color.RGBA{r, g, b, 255})
}

func (i *PNGSaver) SavePNG() bool {
	out, err := os.Create(i.fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return false
	}
	defer out.Close()

	err = png.Encode(out, i.image)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
