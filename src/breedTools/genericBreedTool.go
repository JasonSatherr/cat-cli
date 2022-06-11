package breedTools

import "fmt"

type genericBreedTool struct {
}

func (gbt genericBreedTool) DisplayBreeds() {
	fmt.Print("I don't know what animal breed you want to reveal")
}
