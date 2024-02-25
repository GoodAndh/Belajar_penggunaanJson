package main

import (
	"net/http"
	gacor "webapiok/jauthcon"
)

func main() {
	
	http.HandleFunc("/", gacor.Homepagecan)
	http.ListenAndServe(":3000", nil)

}
