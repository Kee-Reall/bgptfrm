package main

import (
	"blog-api/src/api"
	"log"
	"net/http"
)

func main() {
	r := api.InitMainRouter()

	http.Handle("/", r)

	addr := ":8080"

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	} else {
		log.Println("START SERVICE ON " + addr)
	}

}
