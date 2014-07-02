package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var titles = []string{
	"Star Wars", "The Matrix",
	"Inception", "Hulk",
	"The Departed", "Blade Runner",
	"Alien", "Metropolis",
	"Brazil", "Gattaca",
}

func fetcher(c chan string, api_url string) string {
	resp, err := http.Get(api_url)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	c <- string(content)
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
		//log.Print(len(results))
		elapsed := time.Since(start)
		log.Print("Time elapsed: ", elapsed)
}
