package main

import (
	"fmt"
	"net/http"
)

var PORT = "8000"

func main() {
	// serve the html for swagger UI
	htmlServe := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/", http.StripPrefix("", htmlServe))

	// serve the swagger json files
	swaggerServe := http.FileServer(http.Dir("./swagger"))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", swaggerServe))

	fmt.Println("Swagger UI running on port", PORT, "...")
	http.ListenAndServe(":"+PORT, nil)
}
