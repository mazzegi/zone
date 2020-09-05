package main

import (
	"fmt"

	"github.com/mazzegi/zone"
)

func main() {
	locs, err := zone.Locations()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	for _, loc := range locs {
		off := zone.OffsetToUTC(loc)
		fmt.Printf("Location: %s (%s)\n", loc.String(), off)
	}

	err = zone.Generate("locs.go", "main", "locations")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}
