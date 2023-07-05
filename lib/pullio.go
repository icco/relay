package lib

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Pullio struct {
	Container   string    `json:"container"`
	Image       string    `json:"image"`
	Avatar      string    `json:"avatar"`
	OldImageID  string    `json:"old_image_id"`
	NewImageID  string    `json:"new_image_id"`
	OldVersion  string    `json:"old_version"`
	NewVersion  string    `json:"new_version"`
	OldRevision string    `json:"old_revision"`
	NewRevision string    `json:"new_revision"`
	Type        string    `json:"type"`
	URL         string    `json:"url"`
	Timestamp   time.Time `json:"timestamp"`
}

func jsonToPullio(buf []byte) DataType {
	var data Pullio
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Warnw("decoding json to Pullio", zap.Error(err))
		return nil
	}
	log.Debugw("Pullio data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Pullio) Message() string {
	return fmt.Sprintf("Pullio: updated %q: %q -> %q\n", j.Container, j.OldVersion, j.NewVersion)
}

// Valid checks that the data is good.
func (j *Pullio) Valid() bool {
	return j.Container != "" && j.OldVersion != "" && j.NewVersion != ""
}
