package main

import (
	"github.com/mpsdantas/cvutils"
	"gocv.io/x/gocv"
)

func Draw(image cvutils.Image, color cvutils.Color, point cvutils.Point) {
	image.SetColor(point.X, point.Y, color)
}

func main() {
	name := "bolhas.png"

	// Show image with black trace
	cvutils.ShowIMG(cvutils.ImageOptions{
		Name:       name,
		Flags:      gocv.IMReadGrayScale,
		WindowName: "Bubbles black and white",
		Draw: cvutils.DrawOptions{
			DrawFunc: Draw,
			Color: cvutils.Color{
				Red:   0,
				Green: 0,
				Blue:  0,
			},
			StartingPoint: cvutils.Point{
				X: 200,
				Y: 10,
			},
			EndPoint: cvutils.Point{
				X: 210,
				Y: 200,
			},
		},
	})

	// Show image with red trace
	cvutils.ShowIMG(cvutils.ImageOptions{
		Name:       name,
		Flags:      gocv.IMReadColor,
		WindowName: "Bubbles black and white",
		Draw: cvutils.DrawOptions{
			DrawFunc: Draw,
			Color: cvutils.Color{
				Red:   255,
				Green: 0,
				Blue:  0,
			},
			StartingPoint: cvutils.Point{
				X: 200,
				Y: 10,
			},
			EndPoint: cvutils.Point{
				X: 210,
				Y: 200,
			},
		},
	})
}
