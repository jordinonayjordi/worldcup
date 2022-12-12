/*
Copyright © 2022 Jordi Nonay Jordi <jordinonayjordi@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/jordinonayjordi/worldcup/matches"
	"github.com/spf13/cobra"
)

// TODO
var Today bool
var Tomorrow bool
var Country string

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Get upcoming and past fixtures",
	Long:  `. Get upcoming and past fixtures, including those for your favorite team!`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		games := matches.DownloadMatches()

		if Today {
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

		if Tomorrow {
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

		if Country != "" {
			// TODO check if country in list? check if country inputted
			matchesCountry := games.Country(Country)
			for _, match := range matchesCountry {
				fmt.Printf("\n %s on %s \n %s match at %s \n",
					match.Who(), match.When(), match.Stage(), match.Where())
			}
		}

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
	fixturesCmd.Flags().BoolVarP(&Today, "today", "t", false, "today's fixtures")
	fixturesCmd.Flags().BoolVarP(&Tomorrow, "tomorrow", "w", false, "tomorrow's fixtures")
	fixturesCmd.Flags().StringVarP(&Country, "country", "c", "", "specify country")
}
