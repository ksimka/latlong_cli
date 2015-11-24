package main

import (
	"os"
	"github.com/bradfitz/latlong"
	"fmt"
	"strconv"
)

func main() {
	lat, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		panic("lat")
	}
	lng, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		panic("lng")
	}

	tz := latlong.LookupZoneName(lat, lng)

	fmt.Println(tz)
}