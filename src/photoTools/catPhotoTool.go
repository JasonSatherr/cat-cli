package photoTools

import (
	"fmt"
	"image"
	"net/http"
)

type CatPhotoTool struct {
	//sound string
}

//just move some of this logic out to something that all the classes of the interface can share?
func (cpt CatPhotoTool) GenerateImage() (image.Image, error) {
	fmt.Printf("\n getting an image \n")
	//existingImageFile, err := os.Open("./pics/cat.jpg") //get the cat pic
	existingImageFile, err := http.Get("https://cdn2.thecatapi.com/images/9fm.jpg")
	//https://cdn2.thecatapi.com/images/9fm.jpg
	if err != nil {
		//Bad because we do not unpack the err (*PathError)
		//which contains more information
		return nil, err
	}

	defer existingImageFile.Body.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "jpg"
	imageData, _, err := image.Decode(existingImageFile.Body)

	if err != nil {
		// most likely more we can decode... maybe just pass err back instead of error?
		return nil, err
	}

	// fmt.Println(imageData)
	// fmt.Println(imageType)
	return imageData, nil
}

func (cpt CatPhotoTool) GetExtraMessage() string {
	return "meowww"
	//return cpt.sound
}
