package models

import (
	"time"

	"github.com/scylladb/gocqlx/v3/table"
)

type Monitor struct {
	Uuid      string    `json:"uuid"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func MonitorTableDefinition() table.Metadata {
	return table.Metadata{
		Name:    "monitors",
		Columns: []string{"uuid", "url", "created_at"},
	}
}
