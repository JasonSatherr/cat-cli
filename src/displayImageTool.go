package src

import (
	"image"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/JasonSatherr/cat-cli/src/photoTools"
)

func DisplayCmdHandler(breedId int, animalId string) {
	//determine the url to pull from
	url := determineURL(breedId, animalId)
	//get the image
	img := imageGrabberHandler(breedId, animalId, url)

	//^^ bad because we already did this check, try to change structure of func calls

	//display the image with DisplayRandImage()
	displayRandImage(img)
}
func imageGrabberHandler(breedId int, animalId string, url string) image.Image {
	var img image.Image
	var err error
	if animalId == "cat" {
		catGenerator := photoTools.CatPhotoTool{}
		img, err = catGenerator.GenerateImage()

	}

	if err != nil {
		panic(err)
	}
	return img
}
func determineURL(breedId int, animalId string) string {
	if animalId == "cat" {
		url := "https://api.thecatapi.com/v1/images/search?mime_types=jpg,png"
		if breedId <= 60 && breedId >= 0 {
			url += "/&breed_id=" + strconv.Itoa(breedId)
		}
		return url
	}
	return "https://api.thecatapi.com/v1/images"
}

func displayRandImage(img image.Image) {
	// var Tu float64
	// Tu = -5
	// fmt.Printf("\ntu's number %f", math.Abs(Tu))
	a := app.New()
	w := a.NewWindow("Images")
	displayableImage := canvas.NewImageFromImage(img) //generate image is an image.Image to be rendered...
	w.SetContent(displayableImage)                    //prepare to display image
	w.Resize(fyne.NewSize(640, 480))                  //set size of image
	//lookAtData()
	w.ShowAndRun() //show image
}
