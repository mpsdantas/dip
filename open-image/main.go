package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	name := "biel.png"
	window := gocv.NewWindow("Biel")
	img := gocv.IMRead(name, gocv.IMReadGrayScale)

	defer img.Close()

	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", name)
		return
	}

	window.IMShow(img)
	window.WaitKey(0)
}
