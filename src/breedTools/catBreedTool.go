package breedTools

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/JasonSatherr/cat-cli/src/httpReqTools"
)

type CatBreedTool struct {
	genericBreedTool
}

func (cbt CatBreedTool) DisplayBreeds() {
	cbt.displayCatBreeds()
}

func (cbt CatBreedTool) displayCatBreeds() {
	url := "https://api.thecatapi.com/v1/breeds?attach_breed=0"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "DEMO-API-KEY")

	jsonBytes := httpReqTools.JsonBytesFromRequest(req)
	fmt.Printf("poggers work, length of breed json is... %d\n", len(jsonBytes))
	breedIds := getBreedArrFromBreedJson(jsonBytes)
	fmt.Printf("poggers work, length of breed IDs is... %d\n", len(breedIds))
	for i := 0; i < len(breedIds); i++ {
		fmt.Println(breedIds[i])
	}

}

func getBreedArrFromBreedJson(breedBytes []byte) []string {
	var breedStrings []string

	var jsonTopLevelArr []interface{}
	//turn json bytes into an [] interface{}
	//then for each interface, convert to a map and extracto duh "id"
	err := json.Unmarshal(breedBytes, &jsonTopLevelArr)

	if err != nil {
		panic(err)
	}
	for i := 0; i < len(jsonTopLevelArr); i++ {
		fmt.Println("going through breed")
		breedsMap, ok := jsonTopLevelArr[i].(map[string]interface{}) //map of one breed
		if !ok {
			continue
		}
		id, found := breedsMap["id"]
		if !found {
			fmt.Println("id not found")
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
