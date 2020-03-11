package main

import (
	"fmt"
	"log"

	"github.com/mpsdantas/cvutils"
	"gocv.io/x/gocv"
)

type ImageOptions struct {
	Name       string
	Flags      gocv.IMReadFlag
	WindowName string
	Draw       DrawOptions
}

type DrawOptions struct {
	Color  cvutils.Color
	StartX int
	EndX   int
	StartY int
	EndY   int
}

func ShowIMG(opts ImageOptions) {
	img := cvutils.Image{
		Mat: gocv.IMRead(opts.Name, opts.Flags),
	}

	if img.Mat.Empty() {
		log.Fatal(fmt.Sprintf("Error reading image from: %v\n", opts.Name))
	}

	defer img.Mat.Close()

	window := gocv.NewWindow(opts.WindowName)

	for i := opts.Draw.StartX; i < opts.Draw.EndX; i++ {
		for j := opts.Draw.StartY; j < opts.Draw.EndY; j++ {
			img.SetColor(i, j, opts.Draw.Color)
		}
	}

	window.IMShow(img.Mat)
	window.WaitKey(0)
}

func main() {
	name := "bolhas.png"

	// Show image with black trace
	ShowIMG(ImageOptions{
		Name:       name,
		Flags:      gocv.IMReadGrayScale,
		WindowName: "Bubbles black and white",
		Draw: DrawOptions{
			Color: cvutils.Color{
				Red:   0,
				Green: 0,
				Blue:  0,
			},
			StartX: 200,
			EndX:   210,
			StartY: 10,
			EndY:   200,
		},
	})

	// Show image with red trace
	ShowIMG(ImageOptions{
		Name:       name,
		Flags:      gocv.IMReadColor,
		WindowName: "Bubbles black and white",
		Draw: DrawOptions{
			Color: cvutils.Color{
				Red:   255,
				Green: 0,
				Blue:  0,
			},
			StartX: 200,
			EndX:   210,
			StartY: 10,
			EndY:   200,
		},
	})
}
