package events

import "encoding/json"

type HeartbeatEvent struct {
	MonitorUuid  string `json:"monitor_uuid"`
	StatusCode   uint   `json:"status_code"`
	ResponseTime uint64 `json:"response_time"`
}

func (e HeartbeatEvent) ToString() string {
	json, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}

	return string(json)
}
