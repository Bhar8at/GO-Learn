package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter assembles the response and sends it to the http client
	// http.Request is the client's request and the second parameter in the function contains the path
	// The subslice is used to drop the initial "/"
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])

}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]
	p, _ := LoadPage(title)
	fmt.Fprintf(w, "<h1>%s<h1><p>%s<p>", p.Title, p.Body)

}

func file2() {

	http.HandleFunc("/", viewHandler) // tells the http package to handle all requests to "/"
	log.Fatal(http.ListenAndServe(":8080", nil))
	// ListenAndServe only returns when an unexpected error happens
	// And it will keep listening on port 8080
	// When an error is returned it will be logged by the log.Fatal() fucntion

}
