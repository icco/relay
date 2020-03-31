package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// ReaderToMessage takes in a message buffer and returns a message string.
func ReaderToMessage(b io.Reader) (string, error) {
	var data Data
	var msg string

	buf, err := ioutil.ReadAll(b)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(buf, &data); err != nil {
		return "", fmt.Errorf("decoding json to structured data: %w", err)
	}
	log.WithField("data", data).Debug("data decoded")

	if data.Sonarr != nil {
		for _, ep := range data.Sonarr.Episodes {
			msg += fmt.Sprintf("Sonarr: %s %dx%02d - %q.", data.Sonarr.Series.Title, ep.SeasonNumber, ep.EpisodeNumber, *data.EventType)
		}
	}

	if data.Lidarr != nil {
		for _, ep := range data.Lidarr.Albums {
			msg += fmt.Sprintf("Lidarr: %s - %q - %s", data.Lidarr.Artist.Name, ep.Title, *data.EventType)
		}
	}

	if data.GoogleCloud != nil {
		i := data.GoogleCloud.Incident
		msg += fmt.Sprintf("GCP Alert - %q", i.Summary)
	}

	if msg == "" {
		var f map[string]string
		if err := json.Unmarshal(buf, &f); err != nil {
			return "", fmt.Errorf("decoding json to map: %w", err)
		}

		for k, v := range f {
			msg += fmt.Sprintf("%s: %s\n", k, v)
		}
	}

	return msg, nil
}

// Data is all the types of structured data we can decode.
type Data struct {
	*GoogleCloud
	*Lidarr
	*Sonarr
	EventType *string `json:"eventType"`
}

// Sonarr is the structure of messages we get from Sonarr.
//
// Generated by https://mholt.github.io/json-to-go/
type Sonarr struct {
	Episodes []struct {
		ID             int    `json:"id"`
		EpisodeNumber  int    `json:"episodeNumber"`
		SeasonNumber   int    `json:"seasonNumber"`
		Title          string `json:"title"`
		QualityVersion int    `json:"qualityVersion"`
	} `json:"episodes"`
	Series struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Path   string `json:"path"`
		TvdbID int    `json:"tvdbId"`
	} `json:"series"`
}

// GoogleCloud is the structure of messages we get from Google Cloud Platform Alerting.
//
// Generated by https://mholt.github.io/json-to-go/
type GoogleCloud struct {
	Incident struct {
		IncidentID   string `json:"incident_id"`
		ResourceID   string `json:"resource_id"`
		ResourceName string `json:"resource_name"`
		Resource     struct {
			Type   string `json:"type"`
			Labels struct {
				Host string `json:"host"`
			} `json:"labels"`
		} `json:"resource"`
		ResourceTypeDisplayName string `json:"resource_type_display_name"`
		Metric                  struct {
			Type        string `json:"type"`
			DisplayName string `json:"displayName"`
		} `json:"metric"`
		StartedAt     int    `json:"started_at"`
		PolicyName    string `json:"policy_name"`
		ConditionName string `json:"condition_name"`
		Condition     struct {
			Name               string `json:"name"`
			DisplayName        string `json:"displayName"`
			ConditionThreshold struct {
				Filter       string `json:"filter"`
				Aggregations []struct {
					AlignmentPeriod    string   `json:"alignmentPeriod"`
					PerSeriesAligner   string   `json:"perSeriesAligner"`
					CrossSeriesReducer string   `json:"crossSeriesReducer"`
					GroupByFields      []string `json:"groupByFields"`
				} `json:"aggregations"`
				Comparison     string  `json:"comparison"`
				ThresholdValue float64 `json:"thresholdValue"`
				Duration       string  `json:"duration"`
				Trigger        struct {
					Count int `json:"count"`
				} `json:"trigger"`
			} `json:"conditionThreshold"`
		} `json:"condition"`
		URL     string      `json:"url"`
		State   string      `json:"state"`
		EndedAt interface{} `json:"ended_at"`
		Summary string      `json:"summary"`
	} `json:"incident"`
	Version string `json:"version"`
}

// Lidarr provides a structure for Lidarr updates.
//
// Generated by https://mholt.github.io/json-to-go/
type Lidarr struct {
	Albums []struct {
		ID             int    `json:"id"`
		Title          string `json:"title"`
		QualityVersion int    `json:"qualityVersion"`
	} `json:"albums"`
	Artist struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Path string `json:"path"`
		MbID string `json:"mbId"`
	} `json:"artist"`
}
