package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new command
	// cmd := command
	// // Create a new command handler
	// handler := commandHandler{"Create"}
	// // Create a new command dispatcher
	// dispatcher := CommandDispatcher{}
	// // Add the command handler to the dispatcher
	// dispatcher.AddHandler(handler)
	// // Execute the command
	// dispatcher.Dispatch(cmd)

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello World! madafaka")) })

	log.Fatal(http.ListenAndServe(":8080", router))

}
