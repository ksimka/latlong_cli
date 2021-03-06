package main

import (
	"github.com/ksimka/latlong"
	"fmt"
	"strconv"
	"net/http"
	"os"
	"log"
)

const VERSION = "0.2.0"

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
	if len(os.Args) != 2 {
		log.Fatal("tzname_http [port]")
	}
	if os.Args[1] == "version" {
		fmt.Println(VERSION)
		return
	}

	port := os.Args[1]


	http.HandleFunc("/", tzname)
	http.ListenAndServe(":" + port, nil)
}