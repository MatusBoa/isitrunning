package main

import (
	"isitrunning/visitor/events/kafka"
	"isitrunning/visitor/jobs"

	"github.com/robfig/cron/v3"
)


func main() {
	eventDispatcher := kafka.Create("localhost:9092")
	c := cron.New(cron.WithSeconds())

	c.AddJob("*/30 * * * * *", jobs.RequestPagesJob{
		EventDispatcher: eventDispatcher,
	})

    c.Start()
	defer c.Stop()

	// // Run indefinitely
	select {}
}
