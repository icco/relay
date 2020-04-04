package lib

import (
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
			Want: "Lidarr: Test Name - \"Test title\" - Test\n",
		},
		"plex TV Episode": {
			Have: `{"event":"media.resume","user":true,"owner":true,"Account":{"id":1,"thumb":"https://plex.tv/users/b921c1a7580bf543/avatar?c=1575240953","title":"icco"},"Server":{"title":"storm","uuid":"544b62f0b4f85d5d8f2c91696763d13578f5264a"},"Player":{"local":true,"publicAddress":"68.194.92.253","title":"BRAVIA 4K GB","uuid":"fb39bc1decfcc7a4-com-plexapp-android"},"Metadata":{"librarySectionType":"show","ratingKey":"56914","key":"/library/metadata/56914","parentRatingKey":"56851","grandparentRatingKey":"56850","guid":"com.plexapp.agents.thetvdb://301376/1/14?lang=en","parentGuid":"com.plexapp.agents.thetvdb://301376/1?lang=en","grandparentGuid":"com.plexapp.agents.thetvdb://301376?lang=en","type":"episode","title":"If You Can't Scooby-Doo the Time, Don't Scooby-Doo the Crime","grandparentTitle":"Be Cool, Scooby-Doo!","parentTitle":"Season 1","contentRating":"TV-G","summary":"Fred visits a high security prison and ends up investigating the ghost of an escaped inmate who's wreaking havoc on it.","index":14,"parentIndex":1,"rating":10.0,"lastViewedAt":1585962956,"year":2015,"thumb":"/library/metadata/56914/thumb/1585369990","art":"/library/metadata/56850/art/1585369991","parentThumb":"/library/metadata/56851/thumb/1585369991","grandparentThumb":"/library/metadata/56850/thumb/1585369991","grandparentArt":"/library/metadata/56850/art/1585369991","grandparentTheme":"/library/metadata/56850/theme/1585369991","originallyAvailableAt":"2015-10-31","addedAt":1585287583,"updatedAt":1585369990}}`,
			Want: "Plex - \"media.resume\" : Be Cool, Scooby-Doo! 1x14\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := BufferToMessage([]byte(tc.Have))
			if err != nil {
				t.Error(err)
			}

			if got != tc.Want {
				t.Errorf("BufferToMessage(%q) returned %q wanted %q", tc.Have, got, tc.Want)
			}
		})
	}
}
