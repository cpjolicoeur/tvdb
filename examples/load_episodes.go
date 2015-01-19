package main

import (
	"fmt"
	"github.com/cpjolicoeur/tvdb"
	"log"
)

func main() {
	matches, err := tvdb.GetSeries("A-Team")
	if err != nil {
		log.Fatalln("Error retrieving series results", err.Error())
	}

	if len(matches) < 1 {
		fmt.Println("No series found.")
		return
	}

	series := matches[0]
	fmt.Println("Loading episodes for:", series.Name)
	err = series.GetEpisodes()
	if err != nil {
		log.Fatalln("Error retrieving series episodes", err.Error())
	}

	fmt.Printf("%d total episodes found\n", len(series.Episodes))
	for _, e := range series.Episodes {
		fmt.Printf("s%02.0fe%02.0f - %s\n\t%s\n", e.Season, e.Number, e.Name, e.Overview)
	}
}
