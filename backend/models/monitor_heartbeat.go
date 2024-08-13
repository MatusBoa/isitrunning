package models

import (
	"time"

	"github.com/scylladb/gocqlx/v3/table"
)

type MonitorHeartbeat struct {
	Uuid         string    `json:"uuid"`
	MonitorUuid  string    `json:"monitor_uuid"`
	StatusCode   uint      `json:"status_code"`
	ResponseTime uint64    `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
}

func MonitorHeartbeatTableDefinition() table.Metadata {
	return table.Metadata{
		Name:    "monitor_heartbeats",
		Columns: []string{"uuid", "monitor_uuid", "status_code", "response_time", "created_at"},
	}
}
