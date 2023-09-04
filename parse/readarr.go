package parse

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Readarr struct {
	Author struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		GoodreadsID string `json:"goodreadsId"`
	} `json:"author"`
	Books []struct {
		ID          int    `json:"id"`
		GoodreadsID string `json:"goodreadsId"`
		Title       string `json:"title"`
		Edition     struct {
			Title       string `json:"title"`
			GoodreadsID string `json:"goodreadsId"`
			Asin        string `json:"asin"`
		} `json:"edition"`
		ReleaseDate time.Time `json:"releaseDate"`
	} `json:"books"`
	Release struct {
		Quality           string `json:"quality"`
		QualityVersion    int    `json:"qualityVersion"`
		ReleaseTitle      string `json:"releaseTitle"`
		Indexer           string `json:"indexer"`
		Size              int    `json:"size"`
		CustomFormatScore int    `json:"customFormatScore"`
		//	CustomFormats     []any  `json:"customFormats"`
	} `json:"release"`
	DownloadClient     string `json:"downloadClient"`
	DownloadClientType string `json:"downloadClientType"`
	DownloadID         string `json:"downloadId"`
	EventType          string `json:"eventType"`
	InstanceName       string `json:"instanceName"`
}

func jsonToReadarr(buf []byte) DataType {
	var data Readarr
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Debugw("decoding json to Readarr", zap.Error(err))
		return nil
	}
	log.Debugw("Readarr data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Readarr) Message() string {
	var msg string
	for _, ep := range j.Books {
		msg += fmt.Sprintf("Readarr: %s - %q - %s\n", j.Author.Name, ep.Title, j.EventType)
	}
	return msg
}

// Valid checks that the data is good.
func (j *Readarr) Valid() bool {
	return j.EventType != "" && len(j.Books) > 0 && j.Author.Name != ""
}
