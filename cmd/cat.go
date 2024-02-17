/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
  Short: "Prints cat ASCII art",
  Long:  `This subcommand prints an ASCII art of a cat.`,
  Run: func(cmd *cobra.Command, args []string) {
    // 猫のアスキーアートを出力
    cat := `
/\_/\
( o.o )
> ^ <
    `
    fmt.Println(cat)
  },
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
