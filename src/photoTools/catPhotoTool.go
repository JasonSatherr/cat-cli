package photoTools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type CatPhotoTool struct {
	//sound string
	RandomPhotoTool
}

//just move some of this logic out to something that all the classes of the interface can share?
func (cpt CatPhotoTool) GenerateImageFromUrlEndpoint(url string) (image.Image, error) {
	//existingImageFile, err := os.Open("./pics/cat.jpg") //get the cat pic
	randomCatPicURL, err := cpt.getImgURL(url)
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

func (cpt CatPhotoTool) getImgURL(url string) (string, error) {

	fmt.Println(cpt.getKey())
	//^^need to remove hard link so that we use the img ur
	/*	url := "https://api.thecatapi.com/v1/images/search?size=small&breed_id=2"

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("x-api-key", "DEMO-API-KEY")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)*/

	// response, _ := http.Get(url) //later query specifically for the url??
	// body := response.Body
	// bodyBytes := make([]byte, 0) //The slices to hold all of the body once we buffer it in
	// bodyByteBufferSize := 1024
	// bodyByteBuffer := make([]byte, bodyByteBufferSize) //the buffer to slowly consume the response body
	// var numRead int = bodyByteBufferSize

	// for numRead != 0 { //while there is still data to read...
	// 	read, err := body.Read(bodyByteBuffer)
	// 	numRead = read
	// 	//process duh data
	// 	bodyBytes = append(bodyBytes, bodyByteBuffer[:numRead]...) //bodybytes += bodybytes += buffer

	// 	if err != nil { //print the errors we run into... EOF is largely expected, but others may be present
	// 		fmt.Print(err)
	// 	}
	// }
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "DEMO-API-KEY")
	res, _ := http.DefaultClient.Do(req)
	fmt.Print(url)
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	var jsonData []interface{}

	//get data out of bytes arr and parse it to the form of an array of interfaces
	json.Unmarshal(bodyBytes, &jsonData)

	//because we know the form of the json, we are able to get the image url out of it :)
	if mapData, ok := jsonData[0].(map[string]interface{}); ok {
		if stringToReturn, ok := mapData["url"].(string); ok {
			return stringToReturn, nil
		}
	}
	return "ERROR", errors.New("failed to get the image url :(")
}

func (cpt CatPhotoTool) getKey() string {
	exPath := cpt.getExPath()
	configFileName := exPath + "\\configuration\\secret_config.json"
	//^^IMPORTANT MAKE SURE THAT IT WORKS ON LINUX/MAC TOO NOT JUST WINDOWS
	//^^THIS WAY OF CONCATING TO FILE PATH IS BAD BC OS USE DIFFERENT SYMBOLS TO SEPARATE

	jsonFile, err := os.Open(configFileName)

	if err != nil {
		panic(err)
	}
	byteArr := getByteArrFromFile(jsonFile)

	var jsonData map[string]interface{}

	err = json.Unmarshal(byteArr, &jsonData) //populat jsonData

	if err != nil {
		panic(err)
	}

	keysJson, ok := jsonData["keys"].(map[string]interface{}) //dive further into the json jsonData["keys"]...
	if !ok {
		panic(ok)
	}

	theKey := keysJson["cat-api"].(string) //make the value here equal to a string...

	return theKey
}

func getByteArrFromFile(file *os.File) []byte {
	var byteBuf bytes.Buffer
	lenRead, err := byteBuf.ReadFrom(file)

	if err != nil {
		panic(err)
	}

	fmt.Printf("we read %d bytes", lenRead)

	return byteBuf.Bytes()
}

// func readFileToByteArr(file *os.File, arr []byte) (int32, error){
// 	cumRead := 0
// 	stepSize := 1024
// 	readInStep := 1

// 	for readInStep > 0{
// 		file.Read(arr)

// 	}

// 	return 0, nil
// }

func (cpt CatPhotoTool) getExPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
func (cpt CatPhotoTool) GetExtraMessage() string {
	return "meowww"
}
