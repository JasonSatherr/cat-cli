package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/JasonSatherr/cat-cli/src/photoTools"
)

func main() {
	fmt.Printf("wheeeheee %d", 10)
	// var Tu float64
	// Tu = -5
	// fmt.Printf("\ntu's number %f", math.Abs(Tu))

	catGenerator := photoTools.CatPhotoTool{}
	a := app.New()
	w := a.NewWindow("Images")
	img, err := catGenerator.GenerateImage()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	displayableImage := canvas.NewImageFromImage(img) //generate image is an image.Image to be rendered...
	w.SetContent(displayableImage)                    //prepare to display image
	w.Resize(fyne.NewSize(640, 480))                  //set size of image

	w.ShowAndRun() //show image
}
