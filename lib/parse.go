package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

var (
	funcs = []DataParseFunc{
		jsonToSonarr,
		jsonToLidarr,
		jsonToGoogleCloud,
		jsonToGoogleCloudBuild,
		jsonToPlex,
		jsonToInflux,
		jsonToGitLabPush,
	}
)

// DataParseFunc is the generic func type for parsing input.
type DataParseFunc func([]byte) DataType

// DataType is the interface all parsed data must match.
type DataType interface {
	Valid() bool
	Message() string
}

// BufferToMessage takes in a message buffer and returns a message string.
func BufferToMessage(buf []byte) string {
	var msg string
	log.WithField("body", string(buf)).Debug("attempting to parse")

	for _, f := range funcs {
		if data := f(buf); data != nil {
			if data.Valid() {
				msg += data.Message()
			}
		}
	}

	if msg == "" {
		var f map[string]string
		if err := json.Unmarshal(buf, &f); err != nil {
			log.WithError(err).Info("decoding json to map")
			return ""
		}

		var keys []string
		for k := range f {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			msg += fmt.Sprintf("%s: %s\n", k, f[k])
		}
	}

	return msg
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
	EventType string `json:"eventType"`
	Series    struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Path   string `json:"path"`
		TvdbID int    `json:"tvdbId"`
	} `json:"series"`
}

func jsonToSonarr(buf []byte) DataType {
	var data Sonarr
	if err := json.Unmarshal(buf, &data); err != nil {
		log.WithError(err).Error("decoding json to Sonarr")
		return nil
	}
	log.WithField("data", data).Debug("Sonarr data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Sonarr) Message() string {
	var msg string
	for _, ep := range j.Episodes {
		msg += fmt.Sprintf("Sonarr: %s %dx%02d - %q\n", j.Series.Title, ep.SeasonNumber, ep.EpisodeNumber, j.EventType)
	}

	return msg
}

// Valid checks that the data is good.
func (j *Sonarr) Valid() bool {
	return j.EventType != "" && len(j.Episodes) > 0
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

func jsonToGoogleCloud(buf []byte) DataType {
	var data GoogleCloud
	if err := json.Unmarshal(buf, &data); err != nil {
		log.WithError(err).Error("decoding json to GoogleCloud")
		return nil
	}
	log.WithField("data", data).Debug("GoogleCloud data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *GoogleCloud) Message() string {
	return fmt.Sprintf("GCP Alert - %q\n", j.Incident.Summary)
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
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Status    string `json:"status"`
	Source    struct {
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
		log.WithError(err).Error("decoding json to GoogleCloudBuild")
		return nil
	}
	log.WithField("data", data).Debug("GoogleCloudBuild data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *GoogleCloudBuild) Message() string {
	data, err := base64.StdEncoding.DecodeString(j.Msg.Data)
	if err != nil {
		log.WithError(err).WithField("data", j.Msg.Data).Error("could not decode base64 data")
		return ""
	}
	var sub GCBBuildResource
	if err := json.Unmarshal(data, &sub); err != nil {
		log.WithError(err).Error("decoding json to GoogleCloudBuild Sub")
		return ""
	}

	target := ""
	if len(sub.Artifacts.Images) == 1 {
		target = sub.Artifacts.Images[0]
	}

	return fmt.Sprintf("Google Cloud Build: %s %s @ %s", strings.Title(sub.Status), target, sub.LogURL)
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
		log.WithError(err).Error("decoding json to Lidarr")
		return nil
	}
	log.WithField("data", data).Debug("Lidarr data decoded")

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

// Plex provides a structure for Plex updates.
//
// Generated by https://mholt.github.io/json-to-go/
type Plex struct {
	Event   string `json:"event"`
	User    bool   `json:"user"`
	Owner   bool   `json:"owner"`
	Account struct {
		ID    int    `json:"id"`
		Thumb string `json:"thumb"`
		Title string `json:"title"`
	} `json:"Account"`
	Server struct {
		Title string `json:"title"`
		UUID  string `json:"uuid"`
	} `json:"Server"`
	Player struct {
		Local         bool   `json:"local"`
		PublicAddress string `json:"publicAddress"`
		Title         string `json:"title"`
		UUID          string `json:"uuid"`
	} `json:"Player"`
	Metadata struct {
		AddedAt               int     `json:"addedAt"`
		Art                   string  `json:"art"`
		AudienceRating        float64 `json:"audienceRating"`
		AudienceRatingImage   string  `json:"audienceRatingImage"`
		ChapterSource         string  `json:"chapterSource"`
		ChildCount            int     `json:"childCount"`
		ContentRating         string  `json:"contentRating"`
		Duration              int     `json:"duration"`
		GUID                  string  `json:"guid"`
		GrandparentArt        string  `json:"grandparentArt"`
		GrandparentGUID       string  `json:"grandparentGuid"`
		GrandparentRatingKey  string  `json:"grandparentRatingKey"`
		GrandparentTheme      string  `json:"grandparentTheme"`
		GrandparentThumb      string  `json:"grandparentThumb"`
		GrandparentTitle      string  `json:"grandparentTitle"`
		Index                 int     `json:"index"`
		Key                   string  `json:"key"`
		LastViewedAt          int     `json:"lastViewedAt"`
		LeafCount             int     `json:"leafCount"`
		LibrarySectionType    string  `json:"librarySectionType"`
		OriginallyAvailableAt string  `json:"originallyAvailableAt"`
		ParentGUID            string  `json:"parentGuid"`
		ParentIndex           int     `json:"parentIndex"`
		ParentRatingKey       string  `json:"parentRatingKey"`
		ParentThumb           string  `json:"parentThumb"`
		ParentTitle           string  `json:"parentTitle"`
		PrimaryExtraKey       string  `json:"primaryExtraKey"`
		Rating                float64 `json:"rating"`
		RatingImage           string  `json:"ratingImage"`
		RatingKey             string  `json:"ratingKey"`
		Studio                string  `json:"studio"`
		Summary               string  `json:"summary"`
		Tagline               string  `json:"tagline"`
		Thumb                 string  `json:"thumb"`
		Title                 string  `json:"title"`
		Type                  string  `json:"type"`
		UpdatedAt             int     `json:"updatedAt"`
		ViewCount             int     `json:"viewCount"`
		ViewedLeafCount       int     `json:"viewedLeafCount"`
		Year                  int     `json:"year"`
		Genre                 []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Genre"`
		Director []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Director"`
		Writer []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Writer"`
		Producer []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Producer"`
		Country []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Country"`
		Collection []struct {
			ID  int    `json:"id"`
			Tag string `json:"tag"`
		} `json:"Collection"`
		Role []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count,omitempty"`
			Role  string `json:"role"`
			Thumb string `json:"thumb,omitempty"`
		} `json:"Role"`
		Similar []struct {
			ID    int    `json:"id"`
			Tag   string `json:"tag"`
			Count int    `json:"count"`
		} `json:"Similar"`
		Location []struct {
			Path string `json:"path"`
		} `json:"Location"`
	} `json:"Metadata"`
}

func jsonToPlex(buf []byte) DataType {
	var data Plex
	if err := json.Unmarshal(buf, &data); err != nil {
		log.WithError(err).Error("decoding json to Plex")
		return nil
	}
	log.WithField("data", data).Debug("Plex data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Plex) Message() string {
	switch j.Metadata.Type {
	case "episode":
		return fmt.Sprintf("Plex: %q - %s %dx%d\n", j.Event, j.Metadata.GrandparentTitle, j.Metadata.ParentIndex, j.Metadata.Index)
	case "movie":
		return fmt.Sprintf("Plex: %q - %s\n", j.Event, j.Metadata.Title)
	case "show":
		return fmt.Sprintf("Plex: %q - %s\n", j.Event, j.Metadata.Title)
	default:
		return fmt.Sprintf("Plex: %q - unknown type %q\n", j.Event, j.Metadata.Type)
	}
}

// Valid checks that the data is good.
func (j *Plex) Valid() bool {
	if j.Event == "" {
		return false
	}

	parts := strings.Split(j.Event, ".")
	if len(parts) <= 1 {
		return false
	}

	return parts[0] != "media"
}

// Influx maps the alerts sent from the hosted cloud version of the TICK stack.
//
// Generated by https://mholt.github.io/json-to-go/
type Influx struct {
	CheckID                  string    `json:"_check_id"`
	CheckName                string    `json:"_check_name"`
	Level                    string    `json:"_level"`
	Measurement              string    `json:"_measurement"`
	AlertMessage             string    `json:"_message"`
	NotificationEndpointID   string    `json:"_notification_endpoint_id"`
	NotificationEndpointName string    `json:"_notification_endpoint_name"`
	NotificationRuleID       string    `json:"_notification_rule_id"`
	NotificationRuleName     string    `json:"_notification_rule_name"`
	SourceMeasurement        string    `json:"_source_measurement"`
	SourceTimestamp          int64     `json:"_source_timestamp"`
	Start                    time.Time `json:"_start"`
	StatusTimestamp          int64     `json:"_status_timestamp"`
	Stop                     time.Time `json:"_stop"`
	Time                     time.Time `json:"_time"`
	Type                     string    `json:"_type"`
	Version                  int       `json:"_version"`
}

func jsonToInflux(buf []byte) DataType {
	var data Influx
	if err := json.Unmarshal(buf, &data); err != nil {
		log.WithError(err).Error("decoding json to Influx")
		return nil
	}
	log.WithField("data", data).Debug("Influx data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Influx) Message() string {
	return fmt.Sprintf("TICK Alert: %q", j.AlertMessage)
}

// Valid checks that the data is good.
func (j *Influx) Valid() bool {
	return j.AlertMessage != ""
}

type GitLabPush struct {
	ObjectKind    string `json:"object_kind"`
	EventName     string `json:"event_name"`
	Before        string `json:"before"`
	After         string `json:"after"`
	Ref           string `json:"ref"`
	CheckoutSha   string `json:"checkout_sha"`
	CommitMessage string `json:"message"`
	UserID        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	UserUsername  string `json:"user_username"`
	UserEmail     string `json:"user_email"`
	UserAvatar    string `json:"user_avatar"`
	ProjectID     int    `json:"project_id"`
	Project       struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
		CiConfigPath      string `json:"ci_config_path"`
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		ID        string `json:"id"`
		Message   string `json:"message"`
		Title     string `json:"title"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []string `json:"added"`
		Modified []string `json:"modified"`
		Removed  []string `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
	} `json:"push_options"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

func jsonToGitLabPush(buf []byte) DataType {
	var data GitLabPush
	if err := json.Unmarshal(buf, &data); err != nil {
		log.WithError(err).Error("decoding json to Influx")
		return nil
	}
	log.WithField("data", data).Debug("GitLabPush data decoded")

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *GitLabPush) Message() string {
	commits := ""
	for _, c := range j.Commits {
		commits += fmt.Sprintf(" - %q @ %v %s\n", c.Author.Name, c.Timestamp, c.URL)
	}
	return fmt.Sprintf("GitLab %s to [%s](%s): \n%s", j.EventName, j.Project.Name, j.Project.WebURL, commits)
}

// Valid checks that the data is good.
func (j *GitLabPush) Valid() bool {
	return len(j.Commits) > 0 && j.EventName == "push"
}
