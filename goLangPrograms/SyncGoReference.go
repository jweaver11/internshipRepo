package main

import (
	"fmt"
	"sync"
)

type httpPkg struct{} //Sets 'httpPkg' as a struct

type Container struct {
	mu       sync.Mutex //Mutual exclusion lock that sets 'state' and 'sema' as int
	counters map[string]int
}

var http httpPkg //Sets 'http' as a httpPkg variable

func main() {
	c := Container{ //Sets 'c' as a container
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup //Sets 'wg' as a variable WaitGroup
	var urls = []string{  //Sets 'urls' as a string array with the domains listed
		"http://example.com",
		"http://example.org",
	}

	doIncrement := func(name string, n int) { //Sets 'doIncrement' as a function using a string and an int
		for i := 0; i < n; i++ { //Starts for loop
			c.increment(name) //Calls container to user increment function
		}
		wg.Done() //Subtracts one from the counter
	}

	go doIncrement("a", 10000) //Runs the doIncrement at the same time
	go doIncrement("a", 10000) //They all access the same Container, and two access the same counter
	go doIncrement("b", 10000)

	for _, url := range urls { //Sets a variable 'url' in a for loop that goes through the range of 'urls'
		wg.Add(1) //Adds 1 to the counter of 'wg'

		go func(url string) { //Uses anonymous function call
			defer wg.Done() //Waits until url is retrieved, then lowers counter by one

			http.Get(url) //Gets the url from the saved value
		}(url)
	}

	wg.Wait() //Blocks the WaitGroup counter until it is zero

	fmt.Println(c.counters)

}

func (c *Container) increment(name string) { //Locks the mutex before accessing the counters, and unlocks it at the end of the function using defer
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (httpPkg) Get(url string) {} //Defines function 'Get' using a string and the variable 'httpPkg'
