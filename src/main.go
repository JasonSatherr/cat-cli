package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	fmt.Print(catGenerator)
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
	lookAtData()
	//w.ShowAndRun() //show image
}

func lookAtData() {
	response, _ := http.Get("https://api.thecatapi.com/v1/images/search")
	body := response.Body
	bodyBytes := make([]byte, 16)
	bodyByteBufferSize := 1024
	bodyByteBuffer := make([]byte, bodyByteBufferSize)
	var numRead int = bodyByteBufferSize

	for numRead != 0 {
		read, err := body.Read(bodyByteBuffer)
		numRead = read
		//process duh data
		bodyBytes = append(bodyBytes, bodyByteBuffer[:numRead]...)
		if err != nil {
			//fmt.Print("Uhohhh")
			fmt.Print(err)
		}
		//fmt.Print("woompwoomp")
	}
	fmt.Printf("LEN BODY BYTES = %d", len(bodyBytes))
	var jsonData []interface{}
	json.Unmarshal(bodyBytes, &jsonData)
	fmt.Printf("LEN BODY BYTES = %d", len(bodyBytes))
	fmt.Printf(jsonData[0])
	//fmt.Print(jsonData[0])

}
