//Package breedTools provdes a suite for displaying and imforming the user of acceptable breeds of an animal species
package breedTools

import (
	"fmt"
	"net/http"
	"sort"

	"encoding/json"

	"github.com/JasonSatherr/cat-cli/src/httpReqTools"
)

//CatBreedTool is a struct that has all the utility of the genericBreedTool, but tailored to the cat species
type CatBreedTool struct {
	genericBreedTool
}

//DisplayBreeds is a function that will display the breeds of cats
func (cbt CatBreedTool) DisplayBreeds() {
	cbt.displayCatBreeds()
}

//displayCatBreeds is a function that will fetch data from the cats API
//and unmarshal the data to print out all breeds of cats
func (cbt CatBreedTool) displayCatBreeds() {
	//url is the base endpoint to get data from
	url := "https://api.thecatapi.com/v1/breeds?attach_breed=0"

	//req will be the http request we are preparing to make
	req, _ := http.NewRequest("GET", url, nil)

	//key data to header of req
	req.Header.Add("x-api-key", "DEMO-API-KEY")

	//jsonBytes is the bytes arr of the json body of response
	jsonBytes := httpReqTools.JsonBytesFromRequest(req)

	//breedIds is the list breedId strings
	breedIds := getBreedArrFromBreedJson(jsonBytes)
	sort.Strings(breedIds)

	//print the breeds
	for i := 0; i < len(breedIds); i++ {
		fmt.Println(breedIds[i])
	}

}

//getBreedArrFromBreedJson is a helper function that unpacks json from
//api into an array of breeds
func getBreedArrFromBreedJson(breedBytes []byte) []string {
	//breedStrings will hold the breedId strings
	var breedStrings []string

	var jsonTopLevelArr []interface{}
	//turn json bytes into an [] interface{}
	//then for each interface, convert to a map and extracto duh "id"
	err := json.Unmarshal(breedBytes, &jsonTopLevelArr)

	//if unmarhshal returned nil (look up reasons), panic
	if err != nil {
		panic(err)
	}

	// cycle through json and access ids to add breedStrings slice
	for i := 0; i < len(jsonTopLevelArr); i++ {
		breedsMap, ok := jsonTopLevelArr[i].(map[string]interface{}) //map of one breed
		if !ok {
			continue
		}
		id, found := breedsMap["id"]
		if !found {
			continue

		}
		idStr, ok := id.(string)
		if !ok {
			continue
		}
		breedStrings = append(breedStrings, idStr)
	}
	return breedStrings
}
