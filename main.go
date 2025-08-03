package main

import (
	"blog-api/api"
	"log"
	"net/http"
)

func main() {
	r := api.InitMainRouter()

	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
