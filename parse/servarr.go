package parse

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

type Servarr struct {
	EventMessage    string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
	EventType       string `json:"eventType"`
	InstanceName    string `json:"instanceName"`
}

func jsonToServarr(buf []byte) DataType {
	var data Servarr
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Debugw("decoding json to Servarr", zap.Error(err))
		return nil
	}
	log.Debugw("Servarr data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Servarr) Message() string {
	return fmt.Sprintf("%s: %s\n", j.InstanceName, j.EventMessage)
}

// Valid checks that the data is good.
func (j *Servarr) Valid() bool {
	event := j.EventType == "ApplicationUpdate" || j.EventType == "Health"
	return event && j.EventMessage != "" && j.InstanceName != ""
}
