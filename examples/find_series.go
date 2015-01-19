package main

import (
	"fmt"
	"github.com/cpjolicoeur/tvdb"
	"log"
)

func main() {
	matches, err := tvdb.GetSeries("Lost")
	if err != nil {
		log.Fatalln("Error retrieving series results", err.Error())
	}

	if len(matches) > 0 {
		fmt.Println("The following series matched your search:")
		for _, series := range matches {
			fmt.Printf("%12s - %s\n", series.Id, series.Name)
		}
	} else {
		fmt.Println("No series found.")
	}
}
