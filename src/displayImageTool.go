package src

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/JasonSatherr/cat-cli/src/photoTools"
)

func DisplayCmdHandler(breedId string, animalId string) {
	//determine the url to pull from
	url := DetermineURL(breedId, animalId)
	//get the image
	img := imageGrabberHandler(breedId, animalId, url)

	//^^ bad because we already did this check, try to change structure of func calls

	//display the image with DisplayRandImage()
	displayRandImage(img)
}
func imageGrabberHandler(breedId string, animalId string, url string) image.Image {
	var img image.Image
	var err error
	if animalId == "cat" {
		catGenerator := photoTools.CatPhotoTool{}
		img, _, err = catGenerator.GenerateImageFromUrlEndpoint(url)

	}

	if err != nil {
		panic(err)
	}
	return img
}
func DetermineURL(breedId string, animalId string) string {
	if animalId == "cat" {
		url := "https://api.thecatapi.com/v1/images/search?mime_types=jpg,png"
		if breedId != "" {
			url = url + "&breed_id=" + breedId
		}
		//url := "https://api.thecatapi.com/v1/images/search?breed_id=bali"
		//url += "&breed_id=" + breedId
		return url
	}

	//move into animal photo tools
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
