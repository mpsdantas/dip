package main

import (
	"fmt"
	"image"
	"log"

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

func GetImageRegions(mat gocv.Mat, length int) []gocv.Mat {
	var (
		regions    []gocv.Mat
		startPoint = image.Point{
			X: 0,
			Y: 0,
		}
		endPoint = image.Point{
			X: length / 2,
			Y: length / 2,
		}
	)

	for i := 0; i < 4; i++ {
		regions = append(regions, mat.Region(image.Rect(startPoint.X, startPoint.Y,
			endPoint.X, endPoint.Y)))

		if i%2 == 0 {
			startPoint.Y = length / 2
			endPoint.Y = length
		} else {
			startPoint.X = length / 2
			endPoint.X = length
			startPoint.Y = 0
			endPoint.Y = length / 2
		}
	}

	return regions
}

func ChangeRegions(regions []gocv.Mat) gocv.Mat {
	topHorizontal := gocv.NewMat()
	bottomHorizontal := gocv.NewMat()

	gocv.Vconcat(regions[3], regions[2], &topHorizontal)
	gocv.Vconcat(regions[1], regions[0], &bottomHorizontal)

	finalImage := gocv.NewMat()
	gocv.Hconcat(topHorizontal, bottomHorizontal, &finalImage)

	return finalImage
}

func main() {
	name := "biel.png"

	img := cvutils.Image{
		Mat: gocv.IMRead(name, gocv.IMReadGrayScale),
	}

	if img.Mat.Empty() {
		log.Fatal(fmt.Sprintf("Error reading image from: %v\n", name))
	}

	defer img.Mat.Close()

	dimensions := img.Mat.Size()
	length := dimensions[0]

	if dimensions[0] != dimensions[1] {
		log.Fatal("the image needs to be square, the amount of pixels in " +
			"the horizontal has to be equal to the vertical")
	}

	window := gocv.NewWindow("Change regions")
	// Image:
	// | 0 | 1 |
	// | 2 | 3 |
	// First step is to break the regions
	regions := GetImageRegions(img.Mat, length)

	// Change diagonals:
	// | 3 | 2 |
	// | 1 | 0 |
	finalImage := ChangeRegions(regions)

	window.IMShow(finalImage)
	window.WaitKey(0)
}
