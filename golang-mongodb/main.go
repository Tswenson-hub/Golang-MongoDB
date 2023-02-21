package main

import (
	"fmt"
	"log"
	"net/http"

	"tswenson-hub/golang-mongodb/router"
)

// http://localhost:4000/api/movie
func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000")
}
