//Package fileTools will provide all the functions needed in order to give the cat-cli all the things it needs to interface
//with the file system
package fileTools

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

// func saveImageToFile(image image.Image, fileName string) {
// 	fmt.Printf("we are attempting to save the file name as fileName %s", fileName)
// }

// func saveBytesToFile(bytesArr []byte, fileName string) {
// 	fmt.Printf("we are using an array of bytes to write to the file named %s", fileName)
// }

//SaveToFile is a function that will take in a name (string) and use that as the name of the file that holds the data in the byte arr
func SaveToFile(fileName string, data []byte) {
	pathToFile := ".\\pics\\" + fileName
	fileToWriteTo, err := os.Create(pathToFile)
	if err != nil {
		fmt.Println("Somehow, we failed in creating the file...")
		panic(err)
	}

	defer fileToWriteTo.Close()

	_, err = fileToWriteTo.Write(data)
	if err != nil {
		fmt.Println("Apparently, not all of the data was written to the file")
		panic(err)
	}

}

func SaveImageToFile(fileName string, image image.Image, format string) {
	pathToFile := ".\\pics\\" + fileName
	fileToWriteTo, err := os.Create(pathToFile)
	if err != nil {
		fmt.Println("Somehow, we failed in creating the file...")
		panic(err)
	}

	defer fileToWriteTo.Close()

	if format == "png" {
		err = png.Encode(fileToWriteTo, image)
		if err != nil {
			panic(err)
		}
	} else if format == "jpeg" {
		err = jpeg.Encode(fileToWriteTo, image, nil)
		if err != nil {
			panic(err)
		}
	}
	fmt.Print("image downloaded supposedly")
}
