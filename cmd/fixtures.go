/*
Copyright © 2022 Jordi Nonay Jordi <jordinonayjordi@gmail.com>
*/
package cmd

import (
	"fmt"

	"golang.org/x/exp/slices"

	"github.com/jordinonayjordi/worldcup/matches"
	"github.com/spf13/cobra"
)

// global variables (for flags)
var TodayFix bool
var TomorrowFix bool
var CountryFix string

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Get upcoming and past fixtures",
	Long: `	Get general information about today and tomorrow's fixtures, 
	including stage of competition, teams playing, time and location. 
	You can also find out about your favorite team's past and upcoming games!`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		games := matches.DownloadMatches()

		if TodayFix {
			matchesToday := games.Today()
			if matchesToday == nil {
				fmt.Println("There are no matches today :(")
			} else {
				for _, match := range matchesToday {
					fmt.Printf("\n %s on %s \n %s match at %s \n",
						match.Who(), match.When(), match.Stage(), match.Where())
				}
			}
		}

		if TomorrowFix {
			matchesTmrw := games.Tomorrow()
			if matchesTmrw == nil {
				fmt.Println("There are no matches tomorrow :(")
			} else {
				for _, match := range matchesTmrw {
					fmt.Printf("\n %s on %s \n %s match at %s \n",
						match.Who(), match.When(), match.Stage(), match.Where())
				}
			}
		}

		if CountryFix != "" {
			if slices.Contains(matches.Countries, CountryFix) {
				matchesCountry := games.Country(CountryFix)
				for _, match := range matchesCountry {
					fmt.Printf("\n %s on %s \n %s match at %s \n",
						match.Who(), match.When(), match.Stage(), match.Where())
				}
			} else {
				fmt.Println("The country you selected is not in the World Cup :(")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)
	// today, tomorrow and country flags
	fixturesCmd.Flags().BoolVarP(&TodayFix, "today", "t", false, "get today's fixtures")
	fixturesCmd.Flags().BoolVarP(&TomorrowFix, "tomorrow", "w", false, "get tomorrow's fixtures")
	fixturesCmd.Flags().StringVarP(&CountryFix, "country", "c", "", "get fixtures for country (capitalize first letter)")
}
