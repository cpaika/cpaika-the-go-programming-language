package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a goroutine ( I guess just with the go command?)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //what is this doing?
		// answer - its consuming the goroutine results
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//finally a new function
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch.  What is a channel?
		return
	}
	urlSlice := strings.Split(url, ".") //why is this strings and not string? :shrug:
	domain := urlSlice[len(urlSlice)-2]
	domain = strings.Trim(domain, "https://")

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("While reading from url %s: %v", url, err)
	}
	err = ioutil.WriteFile(domain+".html", body, 0644)
	if err != nil {
		ch <- fmt.Sprintf("While writing to file %s: %v", domain, err)
	}
	resp.Body.Close() //Dont leak resources.  Would the OS clean this up?
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %7d, %s", secs, nbytes, url)

}

// gorountines use channels to communicate.  Channels allow you to pass specified types between concurrent functions
// oh wow this concurrency gets complicated.  I wonder why I never had to think about this in Java?  What does Java do to make this transparent?
// goroutines block until their channel is consumed, and likewise for the consuming thread probably
