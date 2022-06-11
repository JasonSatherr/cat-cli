/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	//"fmt"

	"github.com/JasonSatherr/cat-cli/src/breedTools"
	"github.com/spf13/cobra"
)

var ()

// displayCmd represents the display command
var listCatIDsCmd = &cobra.Command{
	Use:   "listCatIDs",
	Short: "lists the cat breed IDs you can use for flags in cmd display --breedID",
	Long:  `This command is too specific and should become more generalized later on, with the distinction happening with flags`,
	Run: func(cmd *cobra.Command, args []string) {
		//src.DisplayRandImage()
		catBreedTool := breedTools.CatBreedTool{}
		catBreedTool.DisplayBreeds()

	},
}

func init() {
	rootCmd.AddCommand(listCatIDsCmd)
}
