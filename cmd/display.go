/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	//"fmt"

	"github.com/JasonSatherr/cat-cli/src"
	"github.com/spf13/cobra"
)

var (
	breedId  int    = -1
	animalId string = "cat"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "displays an image randomly fetched from the internet",
	Long: `This command will query the internet for a random picture.
	atm, you need to close the image before you can move on`,
	Run: func(cmd *cobra.Command, args []string) {
		src.DisplayCmdHandler(breedId, animalId)
		//src.DisplayRandImage()
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	breedHelpMessage := "a number less than 60 denoting breed ID of cat"
	//Todo: find a way to let this flag only apply to cats///
	//^^TODO, FIND ADAPTIBLE NUMBER OF BREEDS
	displayCmd.Flags().IntVarP(&breedId, "breedID", "b", -1, breedHelpMessage)
	displayCmd.Flags().StringVarP(&animalId, "animalID", "a", "cat", "the animal to display.  'cat' is supported atm 'dog' is soon to come :)")

}
