//Package fileTools will provide all the functions needed in order to give the cat-cli all the things it needs to interface
//with the file system
package fileTools

import (
	"fmt"
	"image"
)

func saveImageToFile(image image.Image, fileName string) {
	fmt.Printf("we are attempting to save the file name as fileName %s", fileName)
}

func saveBytesToFile(bytesArr []byte, fileName string) {
	fmt.Printf("we are using an array of bytes to write to the file named %s", fileName)
}
