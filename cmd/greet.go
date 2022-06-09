/*
Copyright © 2022 Jason Sather

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string = "user"
)

// displayCmd represents the display command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "a silly command so that developers can sandbox with flags",
	Long: `This developer is a bonehead so they really need a way to test things
	at a small scale before trying to do anything real with it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("HELLO %s!!", name)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//greetCmd.Flags().StringVarP(&name, "name", "n", "user", "desires the name of the user")
}
