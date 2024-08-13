package repositories

import (
	"isitrunning/backend/models"
	"time"

	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/table"
)

func CreateMonitorHeartbeatRepository(db *gocqlx.Session) MonitorHeartbeatRepository {
	return MonitorHeartbeatRepository{
		db: db,
	}
}

type MonitorHeartbeatRepository struct {
	db *gocqlx.Session
}

func (r *MonitorHeartbeatRepository) Insert(mh models.MonitorHeartbeat) models.MonitorHeartbeat {
	mh.Uuid = uuid.New().String()
	mh.CreatedAt = time.Now()
	table := table.New(models.MonitorHeartbeatTableDefinition())

	q := r.db.Query(table.Insert()).BindStruct(mh)

	if err := q.ExecRelease(); err != nil {
		panic(err)
	}

	return mh
}
