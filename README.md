# worldcup

`worldcup` is CLI tool to stay updated on how the 2022 World Cup is progressing. The tool allows users to get updates on past fixtures and scores as well as follow the results and upcoming fixtures for their favorite countries.

## Installation

```go
go get github.com/jordinonayjordi/worldcup
```

## Usage

### Commands Available

```
worldcup <command>

Commands:
  fixtures   Get general information about upcoming and past fixtures
  scores     Get scores of past fixtures
  config     Change configuration and defaults

Options:
  -h, --help  Show help                                          [boolean]
```
  

#### Commands `fixtures`
Get general information about today and tomorrow's fixtures, including stage of competition, teams playing, time and location. You can also find out about your favorite team's past and upcoming games!

```
Usage: worldcup fixtures [options]

Options:
  -c, --country   get fixtures for a specific country                   [string]
  -h, --help      help for fixtures                                     [boolean]
  -t, --today     get today’s fixtures                                  [boolean]
  -w, --tomorrow  get tomorrow’s fixtures                               [boolean]

Examples:
  worldcup fixtures -c USA
  worldcup fixtures -t
  worldcup fixtures -w
```

#### Commands `scores`
Get the results of past and today's games and find out how your favorite team's is doing!

```
Usage: worldcup scores [options]

Options:
  -c, --country   get scores for a specific country                   [string]
  -h, --help      help for scores                                     [boolean]
  -t, --today     get today’s scores                                  [boolean]

Examples:
  worldcup scores -c Spain
  worldcup fixtures -t
```

## Acknowledgements
I would like to thank Hanbang Wang for his instruction and continued support throughout the CIS 193 course.

## License
MIT © 2022 Jordi Nonay Jordi
