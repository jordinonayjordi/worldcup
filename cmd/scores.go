/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jordinonayjordi/worldcup/matches"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// global variables (for flags)
var Today bool
var Country string

// scoresCmd represents the scores command
var scoresCmd = &cobra.Command{
	Use:   "scores",
	Short: "Get scores of past fixtures",
	Long: `	Get the results of past and today's games 
	and find out how your favorite team's is doing!`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		games := matches.DownloadMatches()

		if Today {
			matchesToday := games.Today()
			if matchesToday == nil {
				fmt.Println("There are no matches today :(")
			} else {
				for _, match := range matchesToday {
					fmt.Printf("\n %s \n %s match \n",
						match.Result(), match.Stage())
				}
			}
		}

		if Country != "" {
			if slices.Contains(matches.Countries, Country) {
				matchesCountry := games.Country(Country)
				for _, match := range matchesCountry {
					fmt.Printf("\n %s \n %s match \n",
						match.Result(), match.Stage())
				}
			} else {
				fmt.Println("The country you selected is not in the World Cup :(")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scoresCmd)
	// today and country flags
	scoresCmd.Flags().BoolVarP(&Today, "today", "t", false, "get today's scores")
	scoresCmd.Flags().StringVarP(&Country, "country", "c", "", "get scores for country (capitalize first letter)")
}
