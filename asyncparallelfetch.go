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
func fetcher(c chan string, api_url string) string {
	resp, err := http.Get(api_url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //this'll be executed before func return
	body, _ := ioutil.ReadAll(resp.Body)
	c <- string(body)
	return <-c
}

func main() {
		start := time.Now()
		c := make(chan string)
		var results []string

		for i := range titles {
			title := titles[i]
			//for every v, spawn an async call
			go func() {
				c <- fetcher(c, string("http://www.omdbapi.com/?r=JSON&s="+url.QueryEscape(title)))
			}()
		}

		for i := 0; i < len(titles); i++ {
			select {
			case result := <-c:
				results = append(results, result)
				//you can add time out case if you want
			}
		}
		//log.Print(len(results))
		elapsed := time.Since(start)
		log.Print("Time elapsed: ", elapsed)
}
