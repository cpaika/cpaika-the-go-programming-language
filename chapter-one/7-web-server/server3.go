package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting the server on port 8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //Why does this take localhost and not a port?
}

// this is a long one to copy unfortunately
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil { //interesting if statement
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
