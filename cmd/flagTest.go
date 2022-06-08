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
var flagTestCmd = &cobra.Command{
	Use:   "flagTest",
	Short: "a silly command so that developers can sandbox with flags",
	Long: `This developer is a bonehead so they really need a way to test things
	at a small scale before trying to do anything real with it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("flag testing");
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
