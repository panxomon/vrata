package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World! madafaka")) })

	// router.Handle("/get", Get("pancho"))

	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("prueba pal git")

}
