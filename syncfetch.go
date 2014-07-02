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
		for _, v := range titles {
			//log.Print(fetcher("http://www.omdbapi.com/?r=JSON&s=" + url.QueryEscape(v)))
			fetcher("http://www.omdbapi.com/?r=JSON&s=" + url.QueryEscape(v))
		}
		elapsed := time.Since(start)
		log.Print("Time elapsed: ", elapsed)
}
