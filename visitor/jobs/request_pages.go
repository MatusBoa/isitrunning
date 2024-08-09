package jobs

import (
	"isitrunning/visitor/events"
	"isitrunning/visitor/events/kafka"
	"log"
	"net/http"
	"strconv"
)

type RequestPagesJob struct {
	EventDispatcher events.EventDispatcher
}

func (job RequestPagesJob) Run() {
	pages := []string{
		"https://simplo.cz",
	}

	for _, page := range pages {
		go func() {
			log.Print("Requesting", page)
			response, err := http.Get(page)

			if (err != nil) {
				log.Fatal("ERROR", err)
				return
			}

			log.Print(page, response.StatusCode)

			job.EventDispatcher.Dispatch("page", &kafka.KafkaEvent{
				Message: strconv.Itoa(response.StatusCode),
			})
		}()
	}
}