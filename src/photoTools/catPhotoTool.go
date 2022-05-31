package photoTools

import (
	"errors"
	"fmt"
	"image"
	"os"
)

type CatPhotoTool struct {
	//sound string
}

func (cpt CatPhotoTool) GenerateImage() (image.Image, error) {
	fmt.Printf("\n getting an image \n")
	existingImageFile, err := os.Open("./pics/cat.jpg") //get the cat pic

	if err != nil {
		//Bad because we do not unpack the err (*PathError)
		//which contains more information
		return nil, errors.New("COULD NOT FIND IMAGE")
	}

	defer existingImageFile.Close()

	// Calling the generic image.Decode() will tell give us the data
	// and type of image it is as a string. We expect "jpg"
	imageData, _, err := image.Decode(existingImageFile)

	if err != nil {
		// most likely more we can decode... maybe just pass err back instead of error?
		return nil, errors.New("COULD NOT DECODE IMAGE")
	}

	// fmt.Println(imageData)
	// fmt.Println(imageType)
	return imageData, nil
}

func (cpt CatPhotoTool) GetExtraMessage() string {
	return "meowww"
	//return cpt.sound
}
