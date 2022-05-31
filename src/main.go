package main

import (
	"fmt"

	"cat-cli/src/photoTools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	//"github.com/JasonSatherr/cat-cli/src/photoTools"
)

func main() {
	fmt.Printf("wheeeheee %d", 10)
	// var Tu float64
	// Tu = -5
	// fmt.Printf("\ntu's number %f", math.Abs(Tu))

	catGenerator := photoTools.CatPhotoTool{}
	a := app.New()
	w := a.NewWindow("Images")
	img, err := catGenerator.generateImage()
	img = canvas.NewImageFromImage(img) //generate image is an image.Image to be rendered...
	w.SetContent(img)
	w.Resize(fyne.NewSize(640, 480))

	w.ShowAndRun()
}
