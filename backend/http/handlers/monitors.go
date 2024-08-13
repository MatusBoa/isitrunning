package handlers

import (
	"fmt"
	"isitrunning/backend/db"
	"isitrunning/backend/models"
	"isitrunning/backend/repositories"
	"net/http"

	"github.com/labstack/echo"
)

type indexMonitorResponse struct {
	Monitor           models.Monitor            `json:"monitor"`
	MonitorHeartbeats []models.MonitorHeartbeat `json:"monitor_heartbeats"`
}

func IndexMonitors(c echo.Context) error {
	d, err := db.Initialize()

	if err != nil {
		return err
	}

	mr := repositories.CreateMonitorRepository(&d)
	mhr := repositories.CreateMonitorHeartbeatRepository(&d)
	monitors, err := mr.GetAll()

	if err != nil {
		return err
	}

	response := []indexMonitorResponse{}

	for _, monitor := range monitors {
		hearbeats, err := mhr.GetLimitedFromMonitor(monitor.Uuid, 27)

		if err != nil {
			fmt.Println(err)
			return err
		}

		response = append(response, indexMonitorResponse{
			Monitor:           monitor,
			MonitorHeartbeats: hearbeats,
		})
	}

	return c.JSON(http.StatusOK, response)
}
