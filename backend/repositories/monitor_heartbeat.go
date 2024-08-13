package repositories

import (
	"isitrunning/backend/models"
	"time"

	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
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
	mh.CreatedAt = time.Now()
	table := table.New(models.MonitorHeartbeatTableDefinition())

	q := r.db.Query(table.Insert()).BindStruct(mh)

	if err := q.ExecRelease(); err != nil {
		panic(err)
	}

	return mh
}

func (r *MonitorHeartbeatRepository) GetLimitedFromMonitor(monitorUuid string, count uint) ([]models.MonitorHeartbeat, error) {
	var monitorHeartbeats []models.MonitorHeartbeat
	table := table.New(models.MonitorHeartbeatTableDefinition())

	q := qb.Select(table.Name()).
		Where(qb.EqLit("monitor_uuid", monitorUuid)).
		Limit(count).
		OrderBy("created_at", qb.ASC).
		Query(*r.db)

	if err := q.Select(&monitorHeartbeats); err != nil {
		return nil, err
	}

	return monitorHeartbeats, nil
}
