//package breedTool will provide tools to work with the breed of the animals
package breedTools

import "fmt"

type genericBreedTool struct {
}

//DisplayBreeds is a function of genericBreedTool that will allow you to
func (gbt genericBreedTool) DisplayBreeds() {
	fmt.Print("I don't know what animal breed you want to reveal")
}
