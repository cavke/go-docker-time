package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>up and running...</h1>")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format("2006-01-02 15:04:05 -0700 MST")
	fmt.Fprintf(w, "<h1>%s</h1>", time)
}
