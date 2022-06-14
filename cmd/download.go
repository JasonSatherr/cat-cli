/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	"fmt"

	//"github.com/JasonSatherr/cat-cli/src"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sorry, Jason has not yet implemented the download feature")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	//here, we should make sure that the noun is the filename

	//and then make sure that the flags

	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
