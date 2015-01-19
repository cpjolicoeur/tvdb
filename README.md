# tvdb.go

tvdb.go is a client API wrapper for [thetvdb.com][1]

## Installation

Just run `go get github.com/cpjolicoeur/tvdb`

### Examples

    package main

    import "github.com/cpjolicoeur/tvdb"

    func main() {
        // matches is an array of tvdb.Series entries that match
        // the series you were searching for
        matches, err := tvdb.GetSeries("A-Team")

        // To populate the episode and other information for a specific
        // series, you run the `GetEpisodes()` function
        err = matches[0].GetEpisodes()
    }


[1]: http://thetvdb.com/
