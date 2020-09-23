package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	google := "http://www.google.com"
	facebook := "http://www.facebook.com"
	golang := "http://golang.org"
	amazon := "http://www.amazon.com"
	stackoverflow := "http://www.stackoverflow.com"

	links := []string{
		google,
		facebook,
		golang,
		amazon,
		stackoverflow,
	}

	c := make(chan string)

	for _, link := range links {
		go checkSite(link, c)
	}

	//infinite loop, wait until channel sends (l)
	// anon function needs a empty parenthesis to invoke
	// child routine needs a
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkSite(link, c)
		}(l)
	}
}

func checkSite(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s + " is Down")
		c <- s
		return
	}
	fmt.Println(s + " is Up")
	c <- s
}
