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
		"cloudbuild empty": {
			Have: `{"message":{"attributes":{"buildId":"4996a732-a195-49dc-95e3-07d843d0e9bc","status":"QUEUED"},"data":"eyJpZCI6IjQ5OTZhNzMyLWExOTUtNDlkYy05NWUzLTA3ZDg0M2QwZTliYyIsInByb2plY3RJZCI6ImljY28tY2xvdWQiLCJzdGF0dXMiOiJRVUVVRUQiLCJzb3VyY2UiOnsic3RvcmFnZVNvdXJjZSI6eyJidWNrZXQiOiI5NDAzODAxNTQ2MjIuY2xvdWRidWlsZC1zb3VyY2UuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwib2JqZWN0IjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYi0wMzE3YWU2Zi1iYTM3LTQ0NzktYmU3NS1iN2Q4YmUxMDhlYmEudGFyLmd6In19LCJzdGVwcyI6W3sibmFtZSI6Imdjci5pby9jbG91ZC1idWlsZGVycy9kb2NrZXIiLCJhcmdzIjpbImJ1aWxkIiwiLXQiLCJnY3IuaW8vaWNjby1jbG91ZC9yZWxheTo1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIiwiLWYiLCJEb2NrZXJmaWxlIiwiLiJdfV0sImNyZWF0ZVRpbWUiOiIyMDIwLTA0LTI0VDIwOjA1OjExLjY3NDQyNDk4N1oiLCJ0aW1lb3V0IjoiNjAwcyIsImltYWdlcyI6WyJnY3IuaW8vaWNjby1jbG91ZC9yZWxheTo1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIl0sImFydGlmYWN0cyI6eyJpbWFnZXMiOlsiZ2NyLmlvL2ljY28tY2xvdWQvcmVsYXk6NWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYiJdfSwibG9nc0J1Y2tldCI6ImdzOi8vOTQwMzgwMTU0NjIyLmNsb3VkYnVpbGQtbG9ncy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzb3VyY2VQcm92ZW5hbmNlIjp7InJlc29sdmVkU3RvcmFnZVNvdXJjZSI6eyJidWNrZXQiOiI5NDAzODAxNTQ2MjIuY2xvdWRidWlsZC1zb3VyY2UuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwib2JqZWN0IjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYi0wMzE3YWU2Zi1iYTM3LTQ0NzktYmU3NS1iN2Q4YmUxMDhlYmEudGFyLmd6IiwiZ2VuZXJhdGlvbiI6IjE1ODc3NTg3MTE0NjgxNDIifX0sImJ1aWxkVHJpZ2dlcklkIjoiZGVmYXVsdC1naXRodWItY2hlY2tzIiwib3B0aW9ucyI6eyJzdWJzdGl0dXRpb25PcHRpb24iOiJBTExPV19MT09TRSIsImxvZ2dpbmciOiJMRUdBQ1kifSwibG9nVXJsIjoiaHR0cHM6Ly9jb25zb2xlLmNsb3VkLmdvb2dsZS5jb20vY2xvdWQtYnVpbGQvYnVpbGRzLzQ5OTZhNzMyLWExOTUtNDlkYy05NWUzLTA3ZDg0M2QwZTliYz9wcm9qZWN0PTk0MDM4MDE1NDYyMiIsInN1YnN0aXR1dGlvbnMiOnsiQlJBTkNIX05BTUUiOiJtYXN0ZXIiLCJDT01NSVRfU0hBIjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYiIsIlJFUE9fTkFNRSI6InJlbGF5IiwiUkVWSVNJT05fSUQiOiI1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIiwiU0hPUlRfU0hBIjoiNWZiZWI2NyJ9LCJ0YWdzIjpbInRyaWdnZXItZGVmYXVsdC1naXRodWItY2hlY2tzIl19","messageId":"1151174314815328","message_id":"1151174314815328","publishTime":"2020-04-24T20:05:13.67Z","publish_time":"2020-04-24T20:05:13.67Z"},"subscription":"projects/icco-cloud/subscriptions/builds"}`,
			Want: "",
		},
		"cloudbuild": {
			Have: `{ "message": { "publish_time": "2020-04-24T20:49:28.905Z", "data": "eyJpZCI6ImY1N2M4YWQzLTRjNmUtNGIzZC04ZmU2LTQ5ZTM2YzgwNmY5MSIsInByb2plY3RJZCI6ImljY28tY2xvdWQiLCJzdGF0dXMiOiJTVUNDRVNTIiwic291cmNlIjp7InN0b3JhZ2VTb3VyY2UiOnsiYnVja2V0IjoiOTQwMzgwMTU0NjIyLmNsb3VkYnVpbGQtc291cmNlLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsIm9iamVjdCI6ImJlNWY2NzBmYWExYTM2NTI2ODA0ZmY5NGQwNjg4NTU4YWI3OTI2Y2YtNzVhMjY3ZmEtNTBlYy00Y2RlLWJjYzEtNTZlZGJmOWVkZTY5LnRhci5neiJ9fSwic3RlcHMiOlt7Im5hbWUiOiJnY3IuaW8vY2xvdWQtYnVpbGRlcnMvZG9ja2VyIiwiYXJncyI6WyJidWlsZCIsIi10IiwiZ2NyLmlvL2ljY28tY2xvdWQvcmVsYXk6YmU1ZjY3MGZhYTFhMzY1MjY4MDRmZjk0ZDA2ODg1NThhYjc5MjZjZiIsIi1mIiwiRG9ja2VyZmlsZSIsIi4iXSwidGltaW5nIjp7InN0YXJ0VGltZSI6IjIwMjAtMDQtMjRUMjA6NDc6NDQuNzIyMDUwODcxWiIsImVuZFRpbWUiOiIyMDIwLTA0LTI0VDIwOjQ4OjU4Ljk2MTUxMzIzMVoifSwicHVsbFRpbWluZyI6eyJzdGFydFRpbWUiOiIyMDIwLTA0LTI0VDIwOjQ3OjQ0LjcyMjA1MDg3MVoiLCJlbmRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0Nzo0NC43Mjk0MDU1NTNaIn0sInN0YXR1cyI6IlNVQ0NFU1MifV0sInJlc3VsdHMiOnsiaW1hZ2VzIjpbeyJuYW1lIjoiZ2NyLmlvL2ljY28tY2xvdWQvcmVsYXk6YmU1ZjY3MGZhYTFhMzY1MjY4MDRmZjk0ZDA2ODg1NThhYjc5MjZjZiIsImRpZ2VzdCI6InNoYTI1Njo0ODE4NmQ3Yjc4ODIzYTVjNjQ1YTE0ZjNkNGJjYTVhZjExMGU3ZmU3NzUyNWE5YWQyOTFlZTc3YTA4ZjU1NGE4IiwicHVzaFRpbWluZyI6eyJzdGFydFRpbWUiOiIyMDIwLTA0LTI0VDIwOjQ4OjU5LjA1MzkwNzM0NloiLCJlbmRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0OToyNy42NjQ4ODA2ODJaIn19XSwiYnVpbGRTdGVwSW1hZ2VzIjpbInNoYTI1NjplYjgzMjk4NzRkZGZjYjI2MGYyODJiNDcxYzgyMDVlMGI5YTEwZjhkNDJjNDVlZmM4YWIzMjIyMWJjZTQzNDAyIl0sImJ1aWxkU3RlcE91dHB1dHMiOlsiIl19LCJjcmVhdGVUaW1lIjoiMjAyMC0wNC0yNFQyMDo0NzozNy42OTgzMTUzMjdaIiwic3RhcnRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0NzozOS4zOTAwODA5NzFaIiwiZmluaXNoVGltZSI6IjIwMjAtMDQtMjRUMjA6NDk6MjguNTQ0MzYxWiIsInRpbWVvdXQiOiI2MDBzIiwiaW1hZ2VzIjpbImdjci5pby9pY2NvLWNsb3VkL3JlbGF5OmJlNWY2NzBmYWExYTM2NTI2ODA0ZmY5NGQwNjg4NTU4YWI3OTI2Y2YiXSwiYXJ0aWZhY3RzIjp7ImltYWdlcyI6WyJnY3IuaW8vaWNjby1jbG91ZC9yZWxheTpiZTVmNjcwZmFhMWEzNjUyNjgwNGZmOTRkMDY4ODU1OGFiNzkyNmNmIl19LCJsb2dzQnVja2V0IjoiZ3M6Ly85NDAzODAxNTQ2MjIuY2xvdWRidWlsZC1sb2dzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInNvdXJjZVByb3ZlbmFuY2UiOnsicmVzb2x2ZWRTdG9yYWdlU291cmNlIjp7ImJ1Y2tldCI6Ijk0MDM4MDE1NDYyMi5jbG91ZGJ1aWxkLXNvdXJjZS5nb29nbGV1c2VyY29udGVudC5jb20iLCJvYmplY3QiOiJiZTVmNjcwZmFhMWEzNjUyNjgwNGZmOTRkMDY4ODU1OGFiNzkyNmNmLTc1YTI2N2ZhLTUwZWMtNGNkZS1iY2MxLTU2ZWRiZjllZGU2OS50YXIuZ3oiLCJnZW5lcmF0aW9uIjoiMTU4Nzc2MTI1NzU1MTg1MSJ9LCJmaWxlSGFzaGVzIjp7ImdzOi8vOTQwMzgwMTU0NjIyLmNsb3VkYnVpbGQtc291cmNlLmdvb2dsZXVzZXJjb250ZW50LmNvbS9iZTVmNjcwZmFhMWEzNjUyNjgwNGZmOTRkMDY4ODU1OGFiNzkyNmNmLTc1YTI2N2ZhLTUwZWMtNGNkZS1iY2MxLTU2ZWRiZjllZGU2OS50YXIuZ3ojMTU4Nzc2MTI1NzU1MTg1MSI6eyJmaWxlSGFzaCI6W3sidHlwZSI6Ik1ENSIsInZhbHVlIjoiOVloQW9NMzduRy9kTEFQRDZMRVF5dz09In1dfX19LCJidWlsZFRyaWdnZXJJZCI6ImRlZmF1bHQtZ2l0aHViLWNoZWNrcyIsIm9wdGlvbnMiOnsic3Vic3RpdHV0aW9uT3B0aW9uIjoiQUxMT1dfTE9PU0UiLCJsb2dnaW5nIjoiTEVHQUNZIn0sImxvZ1VybCI6Imh0dHBzOi8vY29uc29sZS5jbG91ZC5nb29nbGUuY29tL2Nsb3VkLWJ1aWxkL2J1aWxkcy9mNTdjOGFkMy00YzZlLTRiM2QtOGZlNi00OWUzNmM4MDZmOTE/cHJvamVjdD05NDAzODAxNTQ2MjIiLCJzdWJzdGl0dXRpb25zIjp7IkJSQU5DSF9OQU1FIjoibWFzdGVyIiwiQ09NTUlUX1NIQSI6ImJlNWY2NzBmYWExYTM2NTI2ODA0ZmY5NGQwNjg4NTU4YWI3OTI2Y2YiLCJSRVBPX05BTUUiOiJyZWxheSIsIlJFVklTSU9OX0lEIjoiYmU1ZjY3MGZhYTFhMzY1MjY4MDRmZjk0ZDA2ODg1NThhYjc5MjZjZiIsIlNIT1JUX1NIQSI6ImJlNWY2NzAifSwidGFncyI6WyJ0cmlnZ2VyLWRlZmF1bHQtZ2l0aHViLWNoZWNrcyJdLCJ0aW1pbmciOnsiQlVJTEQiOnsic3RhcnRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0Nzo0My45NTUzMTY4NzZaIiwiZW5kVGltZSI6IjIwMjAtMDQtMjRUMjA6NDg6NTkuMDUzODY3NDY0WiJ9LCJGRVRDSFNPVVJDRSI6eyJzdGFydFRpbWUiOiIyMDIwLTA0LTI0VDIwOjQ3OjQwLjE0Mjc0NzEyMVoiLCJlbmRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0Nzo0My45NTUyNjc1MjhaIn0sIlBVU0giOnsic3RhcnRUaW1lIjoiMjAyMC0wNC0yNFQyMDo0ODo1OS4wNTM5MDY1MjdaIiwiZW5kVGltZSI6IjIwMjAtMDQtMjRUMjA6NDk6MjcuNjY0OTE2NzkzWiJ9fX0=", "message_id": "1151222390869458", "attributes": { "buildId": "f57c8ad3-4c6e-4b3d-8fe6-49e36c806f91", "status": "SUCCESS" } }, "subscription": "projects/icco-cloud/subscriptions/builds" }`,
			Want: "Google Cloud Build: SUCCESS gcr.io/icco-cloud/relay:be5f670faa1a36526804ff94d0688558ab7926cf @ https://console.cloud.google.com/cloud-build/builds/f57c8ad3-4c6e-4b3d-8fe6-49e36c806f91?project=940380154622",
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
			Want: "Plex: \"media.resume\" - Be Cool, Scooby-Doo! 1x14\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := BufferToMessage([]byte(tc.Have))
			if got != tc.Want {
				t.Errorf("BufferToMessage(%q) returned %q wanted %q", tc.Have, got, tc.Want)
			}
		})
	}
}
