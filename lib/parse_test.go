package lib

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		Have string
		Want string
	}{
		"cloud": {
			Have: `{"incident": {"incident_id": "0.ll5vihtvirr7","resource_id": "","resource_name": "Uptime Check URL labels {host=walls.natwelch.com}","resource": {"type": "uptime_url","labels": {"host":"walls.natwelch.com"}},"resource_type_display_name": "Uptime Check URL","metric":{"type": "monitoring.googleapis.com/uptime_check/check_passed", "displayName": "Check passed"},"started_at": 1585430537,"policy_name": "HostDown","condition_name": "Fraction of uptime checks passing per host","condition": {"name":"projects/icco-cloud/alertPolicies/10171777218434756406/conditions/10171777218434756547","displayName":"Fraction of uptime checks passing per host","conditionThreshold":{"filter":"metric.type=\"monitoring.googleapis.com/uptime_check/check_passed\" resource.type=\"uptime_url\"","aggregations":[{"alignmentPeriod":"60s","perSeriesAligner":"ALIGN_NEXT_OLDER","crossSeriesReducer":"REDUCE_FRACTION_TRUE","groupByFields":["resource.label.host"]}],"comparison":"COMPARISON_LT","thresholdValue":0.2,"duration":"600s","trigger":{"count":1}}},"url": "https://console.cloud.google.com/monitoring/alerting/incidents/0.ll5vihtvirr7?project=icco-cloud","state": "open","ended_at": null,"summary": "Check passed for Uptime Check URL labels {host=walls.natwelch.com} is below the threshold of 0.2 with a value of 0.000."},"version": "1.2"}`,
			Want: "GCP Alert - \"Check passed for Uptime Check URL labels {host=walls.natwelch.com} is below the threshold of 0.2 with a value of 0.000.\"\n",
		},
		"sonarr": {
			Have: `{
  "episodes": [
    {
      "id": 123,
      "episodeNumber": 1,
      "seasonNumber": 1,
      "title": "Test title",
      "qualityVersion": 0
    }
  ],
  "eventType": "Test",
  "series": {
    "id": 1,
    "title": "Test Title",
    "path": "C:\\testpath",
    "tvdbId": 1234
  }
}`,
			Want: "Sonarr: Test Title 1x01 - \"Test\"\n",
		},
		"simple": {
			Have: `{"test":"bar","hi":"xyz"}`,
			Want: "hi: xyz\ntest: bar\n",
		},
		"lidarr": {
			Have: `{
  "albums": [
    {
      "id": 123,
      "title": "Test title",
      "qualityVersion": 0
    }
  ],
  "eventType": "Test",
  "artist": {
    "id": 1,
    "name": "Test Name",
    "path": "C:\\testpath",
    "mbId": "aaaaa-aaa-aaaa-aaaaaa"
  }
}`,
			Want: `Lidarr: Test Name - "Test title" - Test`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ReaderToMessage(strings.NewReader(tc.Have))
			if err != nil {
				t.Error(err)
			}

			if got != tc.Want {
				t.Errorf("ReaderToMessage(%q) returned %q wanted %q", tc.Have, got, tc.Want)
			}
		})
	}
}
