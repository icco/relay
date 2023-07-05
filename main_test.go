package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHookHandler(t *testing.T) {
	tt := map[string]struct {
		contentType string
		body        string
	}{
		"empty": {
			contentType: "application/x-www-form-urlencoded",
			body:        "",
		},
		"json": {
			body:        `{"test": "test"}`,
			contentType: "application/json",
		},
		"text": {
			body:        "plain/text",
			contentType: "application/json",
		},
		"form": {},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/hook", nil)
			responseRecorder := httptest.NewRecorder()

			hookHandler(responseRecorder, request)

			if responseRecorder.Code != http.StatusOK {
				t.Errorf("Want status '%d', got '%d'", http.StatusOK, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != "" {
				t.Errorf("Wanted empty response, got %q", responseRecorder.Body)
			}
		})
	}
}
