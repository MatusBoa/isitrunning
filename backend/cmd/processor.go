/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"isitrunning/backend/consumer"
)

// processorCmd represents the processor command
var processorCmd = &cobra.Command{
	Use:   "processor",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		consumer.InitializeConsumer(ctx, &consumer.HeartbeatConsumer{
			ConsumerConfig: consumer.ConsumerConfig{
				ServerAddress: []string{"localhost:9092"},
				Topic:         []string{"heartbeat"},
				Group:         "processor",
			},
		})
	},
}

func init() {
	rootCmd.AddCommand(processorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
