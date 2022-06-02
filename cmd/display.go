/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	"fmt"

	"github.com/JasonSatherr/cat-cli/src"
	"github.com/spf13/cobra"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "displays an image randomly fetched from the internet",
	Long: `This command will query the internet for a random picture.
	atm, you need to close the image before you can move on`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("display called")
		src.DisplayRandImage()
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
