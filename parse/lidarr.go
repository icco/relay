package parse

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

// Lidarr provides a structure for Lidarr updates.
//
// Generated by https://mholt.github.io/json-to-go/
type Lidarr struct {
	Albums []struct {
		ID             int    `json:"id"`
		Title          string `json:"title"`
		QualityVersion int    `json:"qualityVersion"`
	} `json:"albums"`
	EventType string `json:"eventType"`
	Artist    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Path string `json:"path"`
		MbID string `json:"mbId"`
	} `json:"artist"`
}

func jsonToLidarr(buf []byte) DataType {
	var data Lidarr
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Debugw("decoding json to Lidarr", zap.Error(err))
		return nil
	}
	log.Debugw("Lidarr data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Lidarr) Message() string {
	var msg string
	for _, ep := range j.Albums {
		msg += fmt.Sprintf("Lidarr: %s - %q - %s\n", j.Artist.Name, ep.Title, j.EventType)
	}
	return msg
}

// Valid checks that the data is good.
func (j *Lidarr) Valid() bool {
	return j.EventType != "" && len(j.Albums) > 0
}