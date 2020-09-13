// server1 is a minimal echo server, yay
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //wheres the w coming from
}
