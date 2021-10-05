package main

import (
	"image"
	"image/color"
	"math"
	// "golang.org/x/tour/pic"
)

type Image struct {
	data [][]uint8
}

func NewImage(dx, dy int) *Image {
	result := make([][]uint8, dy)
	for y := range result {
		result[y] = make([]uint8, dx)

		for x := range result[y] {
			result[y][x] = uint8(math.Abs(float64(dx/2-x)) + math.Abs(float64(dy/2-y)))
		}
	}

	return &Image{result}
}

func (i *Image) ColorModel() color.Model {
	return color.GrayModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{len(i.data), len(i.data)},
	}
}

func (i *Image) At(x, y int) color.Color {
	return color.Gray{i.data[x][y]}
}

func main() {
	// m := NewImage(100, 100)
	// pic.ShowImage(m)
}
