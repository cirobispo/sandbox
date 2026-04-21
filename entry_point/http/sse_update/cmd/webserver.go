// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func webServer() {
	http.HandleFunc("/events", tickHandler)
	http.Handle("/", http.FileServer(http.Dir("client")))

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Done!")
}

func tickHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter on tickHandler!")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC1123))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		case <-r.Context().Done():
			fmt.Println("Client closed connection")
			return
		}
	}
}
