package main

import (
	"github.com/mpsdantas/cvutils"
	"gocv.io/x/gocv"
)

func Draw(image cvutils.Image, color cvutils.Color, point cvutils.Point) {
	imgColor := image.GetColor(point.X, point.Y)
	inverter := cvutils.Color{
		Red:   255 - imgColor.Red,
		Green: 255 - imgColor.Green,
		Blue:  255 - imgColor.Blue,
	}

	image.SetColor(point.X, point.Y, inverter)
}

func main() {
	name := "biel.png"

	// Show image with black trace
	cvutils.ShowIMG(cvutils.ImageOptions{
		Name:       name,
		Flags:      gocv.IMReadColor,
		WindowName: "Show biel with effect negative",
		Draw: cvutils.DrawOptions{
			DrawFunc: Draw,
			Color: cvutils.Color{
				Red:   0,
				Green: 0,
				Blue:  0,
			},
			StartingPoint: cvutils.Point{
				X: 55,
				Y: 80,
			},
			EndPoint: cvutils.Point{
				X: 210,
				Y: 200,
			},
		},
	})
}
