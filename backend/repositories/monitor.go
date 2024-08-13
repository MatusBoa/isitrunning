package repositories

import (
	"isitrunning/backend/models"

	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
	"github.com/scylladb/gocqlx/v3/table"
)

func CreateMonitorRepository(db *gocqlx.Session) MonitorRepository {
	return MonitorRepository{
		db: db,
	}
}

type MonitorRepository struct {
	db *gocqlx.Session
}

func (r *MonitorRepository) GetAll() ([]models.Monitor, error) {
	var monitors []models.Monitor
	table := table.New(models.MonitorTableDefinition())

	q := r.db.Query(table.Select()).BindMap(qb.M{})

	if err := q.SelectRelease(&monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}
