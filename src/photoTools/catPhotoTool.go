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

	"github.com/JasonSatherr/cat-cli/src/fileTools"
)

type CatPhotoTool struct {
	//sound string
	RandomPhotoTool
}

//just move some of this logic out to something that all the classes of the interface can share?
func (cpt CatPhotoTool) GenerateImageFromUrlEndpoint(url string) (image.Image, error) {
	//existingImageFile, err := os.Open("./pics/cat.jpg") //get the cat pic
	cpt.printSnap()
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

	return imageData, nil
}

func (cpt CatPhotoTool) getImgURL(url string) (string, error) {

	fmt.Println(cpt.getKey()) //get's the key for our API... atm we are just using the demo key :)

	req, _ := http.NewRequest("GET", url, nil) //prepare the http rq

	req.Header.Add("x-api-key", "DEMO-API-KEY") //add headers
	res, _ := http.DefaultClient.Do(req)        //send out and catch the request
	fmt.Print(url)
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	var jsonData []interface{}

	//get data out of bytes arr and parse it to the form of an array of interfaces
	json.Unmarshal(bodyBytes, &jsonData)

	//if //we error if we cannot access jsonData[0]
	//we should access the api once more...

	//because we know the form of the json, we are able to get the image url out of it :)
	if mapData, ok := jsonData[0].(map[string]interface{}); ok { //we error if we cannot access jsonData[0]
		if stringToReturn, ok := mapData["url"].(string); ok {
			return stringToReturn, nil
		}
	}
	return "ERROR", errors.New("failed to get the image url :(")
}

func (cpt CatPhotoTool) getKey() string {
	exPath := fileTools.GetExecutionPath()
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

func (cpt CatPhotoTool) GetExtraMessage() string {
	return "meowww"
}
