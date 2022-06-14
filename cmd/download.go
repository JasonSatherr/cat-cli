/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	"fmt"

	"errors"

	"github.com/JasonSatherr/cat-cli/src"
	"github.com/JasonSatherr/cat-cli/src/fileTools"
	"github.com/JasonSatherr/cat-cli/src/photoTools"
	"github.com/spf13/cobra"
)

var (
	fileName = "cat"
)

// displayCmd represents the display command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "downloads an image randomly fetched from the internet",
	Long: `This command will query the internet for a random picture.
	and then by default download it into the current directory`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires only 1 file name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sorry, Jason is trying to implement the download feature")
		cpt := photoTools.CatPhotoTool{}
		url := src.DetermineURL("", "cat")
		image, format, err := cpt.GenerateImageFromUrlEndpoint(url)
		if err != nil {
			panic(err)
		}
		fileTools.SaveImageToFile(args[0]+"."+format, image, format)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	//here, we should make sure that the noun is the filename

	//and then make sure that the flags

	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
