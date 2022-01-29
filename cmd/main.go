package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World! madafaka")) })

	log.Fatal(http.ListenAndServe(":8080", router))

}
