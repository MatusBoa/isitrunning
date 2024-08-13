package models

import (
	"time"

	"github.com/scylladb/gocqlx/v3/table"
)

type Monitor struct {
	Uuid      string
	Url       string
	CreatedAt time.Time
}

func MonitorTableDefinition() table.Metadata {
	return table.Metadata{
		Name:    "monitors",
		Columns: []string{"uuid", "url", "created_at"},
	}
}
