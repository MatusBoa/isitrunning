package main

import (
	"fmt"
	"isitrunning/visitor/producer"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	Producer producer.Producer
)

func requestPage(url string) {
	log.Print("Requesting", url)
	response, err := http.Get(url)

	if (err != nil) {
		fmt.Println("ERROR", err)
		return
	}

	log.Print(url, response.StatusCode)

	Producer.Send("page", strconv.Itoa(response.StatusCode))
}

func processPages(pages []string) {
	for _, page := range pages {
		go requestPage(page)
	}
}

func main() {
	Producer = producer.CreateProducer("localhost:9092")
	defer Producer.Close()

	// @todo: Get pages from database
	pages := []string{
		"https://simplo.cz",
	}

	processPages(pages)

	tick := time.Tick(5 * time.Second)
	for range tick {
		processPages(pages)
	}
}
