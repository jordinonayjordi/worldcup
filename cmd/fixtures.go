/*
Copyright © 2022 Jordi Nonay Jordi <jordinonayjordi@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/jordinonayjordi/worldcup/matches"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Get upcoming and past fixtures",
	Long:  `. Get upcoming and past fixtures, including those for your favorite team!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fixtures called")

		// decoded := matches.DownloadMatches()

		///

		// country fixtures
		// country, _ := cmd.Flags().GetString("country")
		// TODO check if country in list?
		// call function with arg country
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fixturesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixturesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// TODO
	// fixturesCmd.Flags().StringP("today", "t", "", "today's fixtures") //BoolP?
	// fixturesCmd.Flags().StringP("tomorrow", "w", "", "tomorrow's fixtures") //BoolP?
	// fixturesCmd.Flags().StringP("country", "c", "", "specify country")
}
