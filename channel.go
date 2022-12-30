package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"https://github.com/fabpot",
		"https://github.com/andrew",
		"https://github.com/taylorotwell",
		"https://github.com/egoist",
		"https://github.com/HugoGiraudel",
	}

	c := make(chan string)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)   // This tells the waitgroup, that there is now 1 pending operation here
		go checkUrl(link, c, &wg)
	}

    // this function literal (also called 'anonymous function' or 'lambda expression' in other languages)
    // is useful because 'go' needs to prefix a function and we can save some space by not declaring a whole new function for this
	go func() {
		wg.Wait()	// this blocks the goroutine until WaitGroup counter is zero
		close(c)    // Channels need to be closed, otherwise the below loop will go on forever
	}()    // This calls itself

    // this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel
	for msg := range c {
		fmt.Println(msg)
	}
}

func checkUrls(urls []string) {
	
}

func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	_, err := http.Get(url)

	if err != nil {
		c <- "We could not reach:" + url    // pump the result into the channel
	} else {
		c <- "Success reaching the website:" + url    // pump the result into the channel
	}
}