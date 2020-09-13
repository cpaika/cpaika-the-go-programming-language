// server1 is a minimal echo server, yay
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	fmt.Println("Starting the server on port 8080")
	http.HandleFunc("/", handler) //each handler call is served in a seperate goroutine
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //Why does this take localhost and not a port?
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() //🤮 is this really how I'm supposed to deal with threading? manually locking and unlocking?
	//Crazy how much of the Java multithreading model I take for granted
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) //wheres the w coming from
}

//counter echoes the count of calls to this process.  would be fun to integrate with like sqllite
func counter(w http.ResponseWriter, r *http.Request) { //why is one a pointer and the other isnt
	mu.Lock() //Why do we have to lock to read?
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
