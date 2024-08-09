package jobs

import (
	"isitrunning/visitor/events"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HeartbeatJob struct {
	EventDispatcher events.EventDispatcher
}

func (job HeartbeatJob) Run() {
	pages := []string{
		"https://simplo.cz",
		"https://koterba.sk",
		"https://gateway01.simplo.cz",
	}

	for _, page := range pages {
		go func() {
			log.Printf("Sending heartbeat to %s", page)

			requestCreatedAt := time.Now()
			response, err := http.Get(page)

			if err != nil {
				log.Fatal("ERROR", err)
				return
			}

			defer response.Body.Close()
			responseTime := time.Since(requestCreatedAt).Milliseconds()

			parsed, _ := url.Parse(page)

			job.EventDispatcher.Dispatch("heartbeat", events.HeartbeatEvent{
				Hostname:     parsed.Hostname(),
				Url:          page,
				StatusCode:   uint(response.StatusCode),
				ResponseTime: uint64(responseTime),
			})
		}()
	}
}
