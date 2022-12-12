/*
Copyright © 2022 Jordi Nonay Jordi <jordinonayjordi@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "worldcup",
	Short: "A CLI tool to stay updated on how the 2022 World Cup is progressing",
	Long: `	A CLI tool to stay updated on past and upcoming fixtures 
	as well as scores for the 2022 World Cup. Follow your favorite country's 
	performance along the competition and never miss a match!`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
