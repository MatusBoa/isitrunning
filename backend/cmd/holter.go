/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"isitrunning/backend/events/kafka"
	"isitrunning/backend/jobs"
	"log"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

// holterCmd represents the holter command
var holterCmd = &cobra.Command{
	Use:   "holter",
	Short: "Runs holter service, that periodically collects heartbeats from monitors",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting holter service...")

		eventDispatcher := kafka.CreateEventDispatcher("localhost:9092")
		c := cron.New(cron.WithSeconds())

		c.AddJob("*/10 * * * * *", jobs.HeartbeatJob{
			EventDispatcher: eventDispatcher,
		})

		c.Start()
		defer c.Stop()

		log.Println("Holter is running")

		// // Run indefinitely
		select {}
	},
}

func init() {
	rootCmd.AddCommand(holterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// holterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// holterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
