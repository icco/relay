package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

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

func jsonToGoogleCloud(buf []byte) DataType {
	var data GoogleCloud
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Warnw("decoding json to GoogleCloud", zap.Error(err))
		return nil
	}
	log.Debugw("GoogleCloud data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *GoogleCloud) Message() string {
	return fmt.Sprintf("GCP Alert\n - %q\n - %s\n - <%s>", j.Incident.Summary, j.Incident.PolicyName, j.Incident.URL)
}

// Valid checks that the data is good.
func (j *GoogleCloud) Valid() bool {
	return j.Incident.IncidentID != ""
}

// GoogleCloudBuild are build notifications from Google Cloud Build, delivered via PubSub.
//
// Generated by https://mholt.github.io/json-to-go/
type GoogleCloudBuild struct {
	Msg struct {
		Attributes struct {
			BuildID string `json:"buildId"`
			Status  string `json:"status"`
		} `json:"attributes"`
		Data        string    `json:"data"`
		MessageID   string    `json:"message_id"`
		PublishTime time.Time `json:"publish_time"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// GCBBuildResource is the type that's parsed from the Data base64 encoded
// field.
// https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.builds
// defines the field.
//
// Generated by https://mholt.github.io/json-to-go/
type GCBBuildResource struct {
	ID           string `json:"id"`
	ProjectID    string `json:"projectId"`
	Status       string `json:"status"`
	StatusDetail string `json:"statusDetail"`
	Source       struct {
		StorageSource struct {
			Bucket string `json:"bucket"`
			Object string `json:"object"`
		} `json:"storageSource"`
	} `json:"source"`
	Steps []struct {
		Name string   `json:"name"`
		Args []string `json:"args"`
	} `json:"steps"`
	CreateTime time.Time `json:"createTime"`
	Timeout    string    `json:"timeout"`
	Images     []string  `json:"images"`
	Artifacts  struct {
		Images []string `json:"images"`
	} `json:"artifacts"`
	LogsBucket       string `json:"logsBucket"`
	SourceProvenance struct {
		ResolvedStorageSource struct {
			Bucket     string `json:"bucket"`
			Object     string `json:"object"`
			Generation string `json:"generation"`
		} `json:"resolvedStorageSource"`
	} `json:"sourceProvenance"`
	BuildTriggerID string `json:"buildTriggerId"`
	Options        struct {
		SubstitutionOption string `json:"substitutionOption"`
		Logging            string `json:"logging"`
	} `json:"options"`
	LogURL        string `json:"logUrl"`
	Substitutions struct {
		BRANCHNAME string `json:"BRANCH_NAME"`
		COMMITSHA  string `json:"COMMIT_SHA"`
		REPONAME   string `json:"REPO_NAME"`
		REVISIONID string `json:"REVISION_ID"`
		SHORTSHA   string `json:"SHORT_SHA"`
	} `json:"substitutions"`
	Tags []string `json:"tags"`
}

func jsonToGoogleCloudBuild(buf []byte) DataType {
	var data GoogleCloudBuild
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Warnw("decoding json to GoogleCloudBuild", zap.Error(err))
		return nil
	}
	log.Debugw("GoogleCloudBuild data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *GoogleCloudBuild) Message() string {
	data, err := base64.StdEncoding.DecodeString(j.Msg.Data)
	if err != nil {
		log.Warnw("could not decode base64 data", zap.Error(err), "data", j.Msg.Data)
		return ""
	}
	var sub GCBBuildResource
	if err := json.Unmarshal(data, &sub); err != nil {
		log.Warnw("decoding json to GoogleCloudBuild Sub", zap.Error(err))
		return ""
	}

	return fmt.Sprintf("GCB: %s %v @ <%s>", sub.Status, sub.Artifacts.Images, sub.LogURL)
}

// Valid checks that the data is good.
func (j *GoogleCloudBuild) Valid() bool {
	// https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.builds#Build.Status
	send := false
	switch j.Msg.Attributes.Status {
	case "SUCCESS":
		send = true
	case "FAILURE":
		send = true
	case "INTERNAL_ERROR":
		send = true
	case "TIMEOUT":
		send = true
	}

	return j.Subscription != "" && j.Msg.Data != "" && send
}

type DeployMessage struct {
	Deployed string `json:"deployed"`
	Image    string `json:"image"`
}

func jsonToDeployMessage(buf []byte) DataType {
	var data DeployMessage
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Warnw("decoding json to DeployMessage", zap.Error(err))
		return nil
	}
	log.Debugw("DeployMessage data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *DeployMessage) Message() string {
	return fmt.Sprintf("Deploy: %q -> %q", j.Deployed, j.Image)
}

// Valid checks that the data is good.
func (j *DeployMessage) Valid() bool {
	return j.Deployed != "" && j.Image != ""
}
