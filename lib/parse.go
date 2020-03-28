package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

func ReaderToMessage(b io.Reader) (string, error) {
	var data Data
	if err := json.NewDecoder(b).Decode(&data); err != nil {
		return "", fmt.Errorf("decoding json: %w", err)
	}

	msg := ""
	if data.GenericFlat != nil {
		for k, v := range *data.GenericFlat {
			msg += fmt.Sprintf("%s: %s\n", k, v)
		}
	}

	if data.Sonarr != nil {
		for _, ep := range data.Sonarr.Episodes {
			msg += fmt.Sprintf(" - %s %2dx%2d now available.", data.Sonarr.Series.Title, ep.SeasonNumber, ep.EpisodeNumber)
		}
	}

	if data.GoogleCloud != nil {
		i := data.GoogleCloud.Incident
		msg += fmt.Sprintf("GCP Alert - %q", i.Summary)
	}

	return msg, nil
}

type Data struct {
	*GenericFlat
	*Sonarr
	*GoogleCloud
}

type GenericFlat map[string]string

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
	EventType string `json:"eventType"`
	Series    struct {
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
