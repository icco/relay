package lib

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

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
		log.Debugw("decoding json to Plex", zap.Error(err))
		return nil
	}
	log.Debugw("Plex data decoded", "data", data)

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
	case "album":
		return fmt.Sprintf("Plex: %q - %s by %s\n", j.Event, j.Metadata.Title, j.Metadata.ParentTitle)
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
