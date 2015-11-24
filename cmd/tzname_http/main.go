package main

import (
	"github.com/bradfitz/latlong"
	"fmt"
	"strconv"
	"net/http"
)

func tzname(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 32)
	if err != nil {
		fmt.Fprint(w, "lat")
		return
	}
	lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 32)
	if err != nil {
		fmt.Fprint(w, "lng")
		return
	}

	fmt.Fprint(w, latlong.LookupZoneName(lat, lng))
}

func main() {
	http.HandleFunc("/", tzname)
	http.ListenAndServe(":8080", nil)
}