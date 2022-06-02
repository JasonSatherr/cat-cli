package photoTools

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"net/http"
)

type CatPhotoTool struct {
	//sound string
	RandomPhotoTool
}

//just move some of this logic out to something that all the classes of the interface can share?
func (cpt CatPhotoTool) GenerateImage() (image.Image, error) {
	fmt.Printf("\n getting an image \n")
	//existingImageFile, err := os.Open("./pics/cat.jpg") //get the cat pic
	randomCatPicURL, err := cpt.getImgURL()
	if err != nil {
		return nil, nil
	}
	existingImageFile, err := http.Get(randomCatPicURL)
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

func (cpt CatPhotoTool) getImgURL() (string, error) {
	response, _ := http.Get("https://api.thecatapi.com/v1/images/search") //later query specifically for the url??
	body := response.Body
	bodyBytes := make([]byte, 0)
	bodyByteBufferSize := 1024
	bodyByteBuffer := make([]byte, bodyByteBufferSize)
	var numRead int = bodyByteBufferSize

	for numRead != 0 {
		read, err := body.Read(bodyByteBuffer)
		numRead = read
		//process duh data
		bodyBytes = append(bodyBytes, bodyByteBuffer[:numRead]...)
		fmt.Print((bodyBytes))
		if err != nil {
			fmt.Print(err)
		}
	}
	fmt.Printf("LEN BODY BYTES = %d", len(bodyBytes))
	var jsonData []interface{}
	json.Unmarshal(bodyBytes, &jsonData)
	fmt.Printf("LEN BODY BYTES = %d", len(bodyBytes))

	if mapData, ok := jsonData[0].(map[string]interface{}); ok {
		fmt.Print(mapData["url"])
		if stringToReturn, ok := mapData["url"].(string); ok {
			return stringToReturn, nil
		}
	}
	return "ERROR", errors.New("failed to get the image url :(")
}

func (cpt CatPhotoTool) GetExtraMessage() string {
	return "meowww"
}
