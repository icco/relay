package lib

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Radarr struct {
	Movie struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Year        int    `json:"year"`
		ReleaseDate string `json:"releaseDate"`
		FolderPath  string `json:"folderPath"`
		TmdbID      int    `json:"tmdbId"`
		ImdbID      string `json:"imdbId"`
		Overview    string `json:"overview"`
	} `json:"movie"`
	RemoteMovie struct {
		TmdbID int    `json:"tmdbId"`
		ImdbID string `json:"imdbId"`
		Title  string `json:"title"`
		Year   int    `json:"year"`
	} `json:"remoteMovie"`
	MovieFile struct {
		ID             int       `json:"id"`
		RelativePath   string    `json:"relativePath"`
		Path           string    `json:"path"`
		Quality        string    `json:"quality"`
		QualityVersion int       `json:"qualityVersion"`
		ReleaseGroup   string    `json:"releaseGroup"`
		SceneName      string    `json:"sceneName"`
		IndexerFlags   string    `json:"indexerFlags"`
		Size           int64     `json:"size"`
		DateAdded      time.Time `json:"dateAdded"`
		MediaInfo      struct {
			AudioChannels         float64  `json:"audioChannels"`
			AudioCodec            string   `json:"audioCodec"`
			AudioLanguages        []string `json:"audioLanguages"`
			Height                int      `json:"height"`
			Width                 int      `json:"width"`
			Subtitles             []string `json:"subtitles"`
			VideoCodec            string   `json:"videoCodec"`
			VideoDynamicRange     string   `json:"videoDynamicRange"`
			VideoDynamicRangeType string   `json:"videoDynamicRangeType"`
		} `json:"mediaInfo"`
	} `json:"movieFile"`
	IsUpgrade          bool   `json:"isUpgrade"`
	DownloadClient     string `json:"downloadClient"`
	DownloadClientType string `json:"downloadClientType"`
	DownloadID         string `json:"downloadId"`
	CustomFormatInfo   struct {
		CustomFormats     []any `json:"customFormats"`
		CustomFormatScore int   `json:"customFormatScore"`
	} `json:"customFormatInfo"`
	Release struct {
		ReleaseTitle string `json:"releaseTitle"`
		Indexer      string `json:"indexer"`
		Size         int64  `json:"size"`
	} `json:"release"`
	EventType      string `json:"eventType"`
	InstanceName   string `json:"instanceName"`
	ApplicationURL string `json:"applicationUrl"`
}

func jsonToRadarr(buf []byte) DataType {
	var data Radarr
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Warnw("decoding json to Radarr", zap.Error(err))
		return nil
	}
	log.Debugw("Radarr data decoded", "data", data)

	return &data
}

// Message returns a string representation of this object for human consumption.
func (j *Radarr) Message() string {
	return fmt.Sprintf("Radarr: %q - %s\n", j.Movie.Title, j.EventType)
}

// Valid checks that the data is good.
func (j *Radarr) Valid() bool {
	return j.EventType != "" && len(j.Books) > 0 && j.Author.Name != ""
}
