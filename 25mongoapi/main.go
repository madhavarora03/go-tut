package main

import (
	"fmt"
	"github.com/madhavarora03/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
