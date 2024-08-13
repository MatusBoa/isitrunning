package jobs

import (
	"isitrunning/backend/db"
	"isitrunning/backend/events"
	"isitrunning/backend/repositories"
	"log"
	"net/http"
	"time"
)

type HeartbeatJob struct {
	EventDispatcher events.EventDispatcher
}

func (job HeartbeatJob) Run() {
	d, err := db.Initialize()

	if err != nil {
		panic(err)
	}

	mr := repositories.CreateMonitorRepository(&d)
	monitors, err := mr.GetAll()

	if err != nil {
		panic(err)
	}

	for _, monitor := range monitors {
		go func() {
			log.Printf("Sending heartbeat to %s", monitor.Uuid)

			requestCreatedAt := time.Now()
			response, err := http.Get(monitor.Url)

			if err != nil {
				log.Fatal("ERROR", err)
				return
			}

			defer response.Body.Close()
			responseTime := time.Since(requestCreatedAt).Milliseconds()

			job.EventDispatcher.Dispatch("heartbeat", events.HeartbeatEvent{
				MonitorUuid:  monitor.Uuid,
				StatusCode:   uint(response.StatusCode),
				ResponseTime: uint64(responseTime),
			})
		}()
	}
}
