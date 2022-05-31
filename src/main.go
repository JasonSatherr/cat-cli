package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	fmt.Printf("wheeeheee %d", 10)
	// var Tu float64
	// Tu = -5
	// fmt.Printf("\ntu's number %f", math.Abs(Tu))
	a := app.New()
	w := a.NewWindow("Images")

	img := canvas.NewImageFromImage(catPhtotoTool.generateImage()) //generate image is an image.Image to be rendered...
	w.SetContent(img)
	w.Resize(fyne.NewSize(640, 480))

	w.ShowAndRun()
}
