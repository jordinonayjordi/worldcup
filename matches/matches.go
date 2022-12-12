package matches

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Countries = []string{
	"Qatar", "Ecuador", "Sengal", "Netherlands",
	"England", "Iran", "USA", "Wales",
	"Argentina", "Saudi Arabia", "Mexico", "Poland",
	"Denmark", "Tunisia", "France", "Australia",
	"Croatia", "Belgium", "Canada", "Morocco",
	"Japan", "Spain", "Costa Rica", "Germany",
	"Ghana", "Uruguay", "Korea Republic", "Portugal",
	"Serbia", "Brazil", "Cameroon", "Switzerland",
}

type Match struct {
	MatchNumber   int    `json:"MatchNumber"`
	RoundNumber   int    `json:"RoundNumber"`
	DateUtc       string `json:"DateUtc"` // time.Time
	Location      string `json:"Location"`
	HomeTeam      string `json:"HomeTeam"`
	AwayTeam      string `json:"AwayTeam"`
	Group         string `json:"Group"`
	HomeTeamScore int    `json:"HomeTeamScore"`
	AwayTeamScore int    `json:"AwayTeamScore"`
}

type Matches []Match

func DownloadMatches() Matches {
	resp, err := http.Get("https://fixturedownload.com/feed/json/fifa-world-cup-2022")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var decoded Matches
	err = json.Unmarshal(b, &decoded)
	if err != nil {
		log.Fatal(err)
	}

	return decoded
}

// match Team Matchup
func (match Match) Who() (s string) {
	var matchup string = match.HomeTeam + " vs " + match.AwayTeam
	return matchup
}

// match Results
func (match Match) Result() (s string) {
	var result string = match.HomeTeam + " " + strconv.Itoa(match.HomeTeamScore) +
		" " + match.AwayTeam + " " + strconv.Itoa(match.AwayTeamScore)
	return result
}

// TODO: match Time
func (match Match) When() (s string) {
	t, _ := time.Parse("2006-01-02 15:04:05Z", match.DateUtc)
	loc, _ := time.LoadLocation("America/New_York")
	// var time string = match.DateUtc[:10] + " at " + match.DateUtc[11:19]
	t = (t.In(loc))
	time := t.Month().String() + ", " + strconv.Itoa(t.Day()) + " at " + t.Format("15:04 EST")
	return time
}

// match Location
func (match Match) Where() (s string) {
	var location string = match.Location
	return location
}

// match Stage (e.g., group stage, final, etc.)
func (match Match) Stage() (s string) {
	var stage string
	switch match.RoundNumber {
	case 1, 2, 3:
		stage = "Group Stage"
	case 4:
		stage = "Round of 16"
	case 5:
		stage = "Quarter Final"
	case 6:
		stage = "Semi Final"
	case 7:
		stage = "Final"
	}
	return stage
}

// getting today's matches
func (matches Matches) Today() Matches {
	var gamesToday []Match
	today := time.Now().Format("2006-01-02")
	for _, match := range matches {
		gameDate := match.DateUtc[:10]
		if gameDate == today {
			gamesToday = append(gamesToday, match)
		}
	}
	return gamesToday
}

// getting tomorrow's matches
func (matches Matches) Tomorrow() Matches {
	var gamesTmrw []Match
	tmrw := time.Now().Add(24 * time.Hour).Format("2006-01-02")
	for _, match := range matches {
		gameDate := match.DateUtc[:10]
		if gameDate == tmrw {
			gamesTmrw = append(gamesTmrw, match)
		}
	}
	return gamesTmrw
}

// getting country's matches
func (matches Matches) Country(c string) Matches {
	var gamesCountry []Match
	for _, match := range matches {
		if (c == match.HomeTeam) || (c == match.AwayTeam) {
			gamesCountry = append(gamesCountry, match)
		}
	}
	return gamesCountry
}
