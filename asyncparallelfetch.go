//Execution Type: (asynchromous / concurrent / non-blocking) + (multicore / parallel)
//Goal: Fetch JSON API response for every movie in titles var
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

var titles = []string{
	"Star Wars", "The Matrix",
	"Inception", "Hulk",
	"The Departed", "Blade Runner",
	"Alien", "Metropolis",
	"Brazil", "Gattaca",
}

func init() {
	numcpu := runtime.NumCPU()
	log.Print("setting MAXPROCS to..", numcpu)
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}
func fetcher(api_url string) string {
	resp, err := http.Get(api_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //this'll be executed before func return
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
		start := time.Now()
		c := make(chan string)
		var results []string

		for i := range titles {
			title := titles[i]
			//for every v, spawn an async call
			go func() {
				c <- fetcher(string("http://www.omdbapi.com/?r=JSON&s="+url.QueryEscape(title)))
			}()
		}
		//timeout := time.After(1000 * time.Millisecond)
		for i := 0; i < len(titles); i++ {
			select {
			case result := <-c:
				results = append(results, result)
				//case <-timeout:
				//log.Print("timed out.")
				//continue
			}
		}
	    log.Print(len(results))
		elapsed := time.Since(start)
		log.Print("Time elapsed: ", elapsed)
}
