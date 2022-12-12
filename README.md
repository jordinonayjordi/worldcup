# worldcup

A CLI tool to stay updated on how the 2022 World Cup is progressing

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
Get general information about upcoming and past fixtures, including stage of competition, teams playing, time and location

```
Usage: worldcup fixtures [options]

Options:
  -h, --help      Show help                                             [boolean]
  -t, --today     Get today’s fixtures                                  [boolean]
  -w, --tomorrow  Get tomorrow’s fixtures                               [boolean]
  -c, --country   Get fixtures for your selected country                [string]

Examples:
  worldcup fixtures -t
  worldcup fixtures -w
  worldcup fixtures -c Spain
```

#### Commands `scores`
Get scores of past fixtures



## Acknowledgements
I would like to thank .....

## License
MIT © 2022 Jordi Nonay Jordi
