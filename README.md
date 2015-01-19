# tvdb.go

tvdb.go is a client API wrapper for [thetvdb.com][1]

## Installation

Just run `go get github.com/cpjolicoeur/tvdb`

### Examples

    package main

    import "github.com/cpjolicoeur/tvdb"

    func main() {
        matches, err := tvdb.GetSeries("A-Team")
    }


[1]: http://thetvdb.com/
