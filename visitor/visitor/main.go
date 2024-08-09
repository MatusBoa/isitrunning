package visitor

import (
	"fmt"
	"isitrunning/visitor/producer"
	"log"
	"net/http"
	"strconv"
)

type Visitor struct {
	Producer producer.Producer
}

func (v Visitor) Run() {
	pages := []string{
		"https://simplo.cz",
	}

	for _, page := range pages {
		go v.requestPage(page)
	}
}

func (v *Visitor) requestPage(url string) {
	log.Print("Requesting", url)
	response, err := http.Get(url)

	if (err != nil) {
		fmt.Println("ERROR", err)
		return
	}

	log.Print(url, response.StatusCode)

	v.Producer.Send("page", strconv.Itoa(response.StatusCode))
}