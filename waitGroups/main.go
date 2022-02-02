package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSize(url string) {
	//SChedule the call to WaitGroup's Done to tell goroutine is complete
	defer wg.Done()

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 2: ", url)
	defer response.Body.Close()

	fmt.Println("Step 3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step 4: ", len(body))
}

func main() {
	// Add a count of three, one for each goroutine
	wg.Add(5)
	fmt.Println("Start Goroutines")

	go responseSize("https://www.golangprograms.com")
	go responseSize("https://stackoverflow.com")
	go responseSize("https://coderwall.com")
	go responseSize("http://lax.dog")
	go responseSize("https://sdlkjaslkdjalksdjla.com")

	wg.Wait()
	fmt.Println("Terminate")
}
