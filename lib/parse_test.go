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
			Want: "GCP Alert\n - \"Check passed for Uptime Check URL labels {host=walls.natwelch.com} is below the threshold of 0.2 with a value of 0.000.\"\n - HostDown\n - <https://console.cloud.google.com/monitoring/alerting/incidents/0.ll5vihtvirr7?project=icco-cloud>",
		},
		"cloudbuild empty": {
			Have: `{"message":{"attributes":{"buildId":"4996a732-a195-49dc-95e3-07d843d0e9bc","status":"QUEUED"},"data":"eyJpZCI6IjQ5OTZhNzMyLWExOTUtNDlkYy05NWUzLTA3ZDg0M2QwZTliYyIsInByb2plY3RJZCI6ImljY28tY2xvdWQiLCJzdGF0dXMiOiJRVUVVRUQiLCJzb3VyY2UiOnsic3RvcmFnZVNvdXJjZSI6eyJidWNrZXQiOiI5NDAzODAxNTQ2MjIuY2xvdWRidWlsZC1zb3VyY2UuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwib2JqZWN0IjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYi0wMzE3YWU2Zi1iYTM3LTQ0NzktYmU3NS1iN2Q4YmUxMDhlYmEudGFyLmd6In19LCJzdGVwcyI6W3sibmFtZSI6Imdjci5pby9jbG91ZC1idWlsZGVycy9kb2NrZXIiLCJhcmdzIjpbImJ1aWxkIiwiLXQiLCJnY3IuaW8vaWNjby1jbG91ZC9yZWxheTo1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIiwiLWYiLCJEb2NrZXJmaWxlIiwiLiJdfV0sImNyZWF0ZVRpbWUiOiIyMDIwLTA0LTI0VDIwOjA1OjExLjY3NDQyNDk4N1oiLCJ0aW1lb3V0IjoiNjAwcyIsImltYWdlcyI6WyJnY3IuaW8vaWNjby1jbG91ZC9yZWxheTo1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIl0sImFydGlmYWN0cyI6eyJpbWFnZXMiOlsiZ2NyLmlvL2ljY28tY2xvdWQvcmVsYXk6NWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYiJdfSwibG9nc0J1Y2tldCI6ImdzOi8vOTQwMzgwMTU0NjIyLmNsb3VkYnVpbGQtbG9ncy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzb3VyY2VQcm92ZW5hbmNlIjp7InJlc29sdmVkU3RvcmFnZVNvdXJjZSI6eyJidWNrZXQiOiI5NDAzODAxNTQ2MjIuY2xvdWRidWlsZC1zb3VyY2UuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwib2JqZWN0IjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYi0wMzE3YWU2Zi1iYTM3LTQ0NzktYmU3NS1iN2Q4YmUxMDhlYmEudGFyLmd6IiwiZ2VuZXJhdGlvbiI6IjE1ODc3NTg3MTE0NjgxNDIifX0sImJ1aWxkVHJpZ2dlcklkIjoiZGVmYXVsdC1naXRodWItY2hlY2tzIiwib3B0aW9ucyI6eyJzdWJzdGl0dXRpb25PcHRpb24iOiJBTExPV19MT09TRSIsImxvZ2dpbmciOiJMRUdBQ1kifSwibG9nVXJsIjoiaHR0cHM6Ly9jb25zb2xlLmNsb3VkLmdvb2dsZS5jb20vY2xvdWQtYnVpbGQvYnVpbGRzLzQ5OTZhNzMyLWExOTUtNDlkYy05NWUzLTA3ZDg0M2QwZTliYz9wcm9qZWN0PTk0MDM4MDE1NDYyMiIsInN1YnN0aXR1dGlvbnMiOnsiQlJBTkNIX05BTUUiOiJtYXN0ZXIiLCJDT01NSVRfU0hBIjoiNWZiZWI2NzZmMWYzYmNiNjIyYWUwMDkyMDBkODQ1N2Y2NWE0MGFmYiIsIlJFUE9fTkFNRSI6InJlbGF5IiwiUkVWSVNJT05fSUQiOiI1ZmJlYjY3NmYxZjNiY2I2MjJhZTAwOTIwMGQ4NDU3ZjY1YTQwYWZiIiwiU0hPUlRfU0hBIjoiNWZiZWI2NyJ9LCJ0YWdzIjpbInRyaWdnZXItZGVmYXVsdC1naXRodWItY2hlY2tzIl19","messageId":"1151174314815328","message_id":"1151174314815328","publishTime":"2020-04-24T20:05:13.67Z","publish_time":"2020-04-24T20:05:13.67Z"},"subscription":"projects/icco-cloud/subscriptions/builds"}`,
			Want: "",
		},
		"cloudbuild": {
			Have: `{"message":{"attributes":{"buildId":"3b34ea43-928e-4b32-95c6-47f47c683d5e","status":"SUCCESS"},"data":"ewogICJpZCI6ICIzYjM0ZWE0My05MjhlLTRiMzItOTVjNi00N2Y0N2M2ODNkNWUiLAogICJzdGF0dXMiOiAiU1VDQ0VTUyIsCiAgInNvdXJjZSI6IHsKICB9LAogICJjcmVhdGVUaW1lIjogIjIwMjItMDQtMTZUMTU6NDA6MTkuNDQyNDQ3NTk3WiIsCiAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQwOjIwLjI4OTg4MDk3OVoiLAogICJmaW5pc2hUaW1lIjogIjIwMjItMDQtMTZUMTU6NDY6MjYuNzQ3NTgwWiIsCiAgInJlc3VsdHMiOiB7CiAgICAiaW1hZ2VzIjogW3sKICAgICAgIm5hbWUiOiAiZ2NyLmlvL2ljY28tY2xvdWQvbG9jYXRpdmUuZ2FyZGVuIiwKICAgICAgImRpZ2VzdCI6ICJzaGEyNTY6Mzk5OGZkN2VmNGU1ODJkYmJjZjAxOTNhZmVkNTUzNzYxMmYzOGE1YzdhMTMxMjExMDIzMmNhNjczNmU0NjM5NSIsCiAgICAgICJwdXNoVGltaW5nIjogewogICAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMy4xOTU5Njk4MDZaIiwKICAgICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ2OjI0LjY0MTUzMzE1MloiCiAgICAgIH0KICAgIH0sIHsKICAgICAgIm5hbWUiOiAiZ2NyLmlvL2ljY28tY2xvdWQvbG9jYXRpdmUuZ2FyZGVuOmxhdGVzdCIsCiAgICAgICJkaWdlc3QiOiAic2hhMjU2OjM5OThmZDdlZjRlNTgyZGJiY2YwMTkzYWZlZDU1Mzc2MTJmMzhhNWM3YTEzMTIxMTAyMzJjYTY3MzZlNDYzOTUiLAogICAgICAicHVzaFRpbWluZyI6IHsKICAgICAgICAic3RhcnRUaW1lIjogIjIwMjItMDQtMTZUMTU6NDY6MjMuMTk1OTY5ODA2WiIsCiAgICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyNC42NDE1MzMxNTJaIgogICAgICB9CiAgICB9LCB7CiAgICAgICJuYW1lIjogImdjci5pby9pY2NvLWNsb3VkL2xvY2F0aXZlLmdhcmRlbjo3ZWY3ZWU2NzUyMTVjY2QwODA3MzEwYzAxZGNjZmNlM2VlMjk1MTBkIiwKICAgICAgImRpZ2VzdCI6ICJzaGEyNTY6Mzk5OGZkN2VmNGU1ODJkYmJjZjAxOTNhZmVkNTUzNzYxMmYzOGE1YzdhMTMxMjExMDIzMmNhNjczNmU0NjM5NSIsCiAgICAgICJwdXNoVGltaW5nIjogewogICAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMy4xOTU5Njk4MDZaIiwKICAgICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ2OjI0LjY0MTUzMzE1MloiCiAgICAgIH0KICAgIH1dLAogICAgImJ1aWxkU3RlcEltYWdlcyI6IFsic2hhMjU2OmQ0ZDNkY2FiZDM2NjI2MzY0YmM4YjNhNDFiMzEwZjNhZjMwYjJjMDFiN2FiMTk3MDc5NDc5ZTM3NDgyZjAxY2EiLCAic2hhMjU2OmQ0ZDNkY2FiZDM2NjI2MzY0YmM4YjNhNDFiMzEwZjNhZjMwYjJjMDFiN2FiMTk3MDc5NDc5ZTM3NDgyZjAxY2EiLCAic2hhMjU2OmQ0ZDNkY2FiZDM2NjI2MzY0YmM4YjNhNDFiMzEwZjNhZjMwYjJjMDFiN2FiMTk3MDc5NDc5ZTM3NDgyZjAxY2EiLCAic2hhMjU2OmVhN2IxODI0NDkyZGYxNmI0MTE0ODkwNDVmM2Q1OWJmZjM2NWY5YWMyYzljODBmMjM4ZTExMTVmZDUwMmMwZTMiLCAic2hhMjU2OmMxYzFjZGE3MmFiOGMzMDYzOTBmYzA1NTE4YmY0ZDQyMTQ4NTY0OTc4MzI2ZDA3OGY2NWQ1NDY4NThkMTM5Y2IiXSwKICAgICJidWlsZFN0ZXBPdXRwdXRzIjogWyIiLCAiIiwgIiIsICIiLCAiIl0KICB9LAogICJzdGVwcyI6IFt7CiAgICAibmFtZSI6ICJnY3IuaW8vY2xvdWQtYnVpbGRlcnMvZG9ja2VyIiwKICAgICJhcmdzIjogWyJidWlsZCIsICItdCIsICJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46N2VmN2VlNjc1MjE1Y2NkMDgwNzMxMGMwMWRjY2ZjZTNlZTI5NTEwZCIsICItdCIsICJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46bGF0ZXN0IiwgIi4iLCAiLWYiLCAiRG9ja2VyZmlsZSJdLAogICAgImlkIjogIkJ1aWxkIiwKICAgICJ0aW1pbmciOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0MDoyNC45MTMyMDgyMTNaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NDoyNy44NzI3OTUyMjZaIgogICAgfSwKICAgICJzdGF0dXMiOiAiU1VDQ0VTUyIsCiAgICAicHVsbFRpbWluZyI6IHsKICAgICAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQwOjI0LjkxMzIwODIxM1oiLAogICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQwOjI0LjkxNzIwMzI4MFoiCiAgICB9CiAgfSwgewogICAgIm5hbWUiOiAiZ2NyLmlvL2Nsb3VkLWJ1aWxkZXJzL2RvY2tlciIsCiAgICAiYXJncyI6IFsicHVzaCIsICJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46N2VmN2VlNjc1MjE1Y2NkMDgwNzMxMGMwMWRjY2ZjZTNlZTI5NTEwZCJdLAogICAgImlkIjogIlB1c2ggU0hBIiwKICAgICJ0aW1pbmciOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NDoyNy44NzM3MDAzMDJaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NDo1Ny43OTgzNzkyODNaIgogICAgfSwKICAgICJzdGF0dXMiOiAiU1VDQ0VTUyIsCiAgICAicHVsbFRpbWluZyI6IHsKICAgICAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ0OjI3Ljg3MzcwMDMwMloiLAogICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ0OjI3Ljg3OTg4MjcxOFoiCiAgICB9CiAgfSwgewogICAgIm5hbWUiOiAiZ2NyLmlvL2Nsb3VkLWJ1aWxkZXJzL2RvY2tlciIsCiAgICAiYXJncyI6IFsicHVzaCIsICJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46bGF0ZXN0Il0sCiAgICAiaWQiOiAiUHVzaCBsYXRlc3QiLAogICAgInRpbWluZyI6IHsKICAgICAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ0OjU3Ljc5ODU1MzIzNloiLAogICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ0OjU5LjQ5NTMyMzMzNloiCiAgICB9LAogICAgInN0YXR1cyI6ICJTVUNDRVNTIiwKICAgICJwdWxsVGltaW5nIjogewogICAgICAic3RhcnRUaW1lIjogIjIwMjItMDQtMTZUMTU6NDQ6NTcuNzk4NTUzMjM2WiIsCiAgICAgICJlbmRUaW1lIjogIjIwMjItMDQtMTZUMTU6NDQ6NTcuODAwMTM4MjUyWiIKICAgIH0KICB9LCB7CiAgICAibmFtZSI6ICJnY3IuaW8vZ29vZ2xlLmNvbS9jbG91ZHNka3Rvb2wvY2xvdWQtc2RrOnNsaW0iLAogICAgImFyZ3MiOiBbInJ1biIsICJzZXJ2aWNlcyIsICJ1cGRhdGUiLCAibG9jYXRpdmUiLCAiLS1wbGF0Zm9ybVx1MDAzZG1hbmFnZWQiLCAiLS1pbWFnZVx1MDAzZGdjci5pby9pY2NvLWNsb3VkL2xvY2F0aXZlLmdhcmRlbjo3ZWY3ZWU2NzUyMTVjY2QwODA3MzEwYzAxZGNjZmNlM2VlMjk1MTBkIiwgIi0tbGFiZWxzXHUwMDNkbWFuYWdlZC1ieVx1MDAzZGdjcC1jbG91ZC1idWlsZC1kZXBsb3ktY2xvdWQtcnVuLGNvbW1pdC1zaGFcdTAwM2Q3ZWY3ZWU2NzUyMTVjY2QwODA3MzEwYzAxZGNjZmNlM2VlMjk1MTBkLGdjYi1idWlsZC1pZFx1MDAzZDNiMzRlYTQzLTkyOGUtNGIzMi05NWM2LTQ3ZjQ3YzY4M2Q1ZSxnY2ItdHJpZ2dlci1pZFx1MDAzZDM3ZTIxZjQzLTI1MzAtNDVjNy1hNTVjLTEwOGJhNDAzMTA0YiIsICItLXJlZ2lvblx1MDAzZHVzLWNlbnRyYWwxIiwgIi0tcXVpZXQiXSwKICAgICJpZCI6ICJEZXBsb3kiLAogICAgImVudHJ5cG9pbnQiOiAiZ2Nsb3VkIiwKICAgICJ0aW1pbmciOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NDo1OS40OTU0MzA5MzdaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMS4yNzc2MTY0NDZaIgogICAgfSwKICAgICJzdGF0dXMiOiAiU1VDQ0VTUyIsCiAgICAicHVsbFRpbWluZyI6IHsKICAgICAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ0OjU5LjQ5NTQzMDkzN1oiLAogICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ1OjU4LjE1MzI2NzI2MVoiCiAgICB9CiAgfSwgewogICAgIm5hbWUiOiAiY3VybGltYWdlcy9jdXJsIiwKICAgICJhcmdzIjogWyItc3ZMIiwgIi1kIiwgIlwie1xcXCJkZXBsb3llZFxcXCI6IFxcXCJsb2NhdGl2ZVxcXCIsIFxcXCJpbWFnZVxcXCI6IFxcXCJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46N2VmN2VlNjc1MjE1Y2NkMDgwNzMxMGMwMWRjY2ZjZTNlZTI5NTEwZFxcXCJ9XCIiLCAiLVgiLCAiUE9TVCIsICItLWhlYWRlciIsICJDb250ZW50LVR5cGU6IGFwcGxpY2F0aW9uL2pzb24iLCAiLWYiLCAiaHR0cHM6Ly9yZWxheS5uYXR3ZWxjaC5jb20vaG9vayJdLAogICAgImlkIjogIk5vdGZpeSIsCiAgICAidGltaW5nIjogewogICAgICAic3RhcnRUaW1lIjogIjIwMjItMDQtMTZUMTU6NDY6MjEuMjc3NzY0NDA5WiIsCiAgICAgICJlbmRUaW1lIjogIjIwMjItMDQtMTZUMTU6NDY6MjMuMDQ1MDQ4ODY2WiIKICAgIH0sCiAgICAic3RhdHVzIjogIlNVQ0NFU1MiLAogICAgInB1bGxUaW1pbmciOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMS4yNzc3NjQ0MDlaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMi4xNTkyNDE4NzdaIgogICAgfQogIH1dLAogICJ0aW1lb3V0IjogIjEyMDBzIiwKICAiaW1hZ2VzIjogWyJnY3IuaW8vaWNjby1jbG91ZC9sb2NhdGl2ZS5nYXJkZW46bGF0ZXN0IiwgImdjci5pby9pY2NvLWNsb3VkL2xvY2F0aXZlLmdhcmRlbjo3ZWY3ZWU2NzUyMTVjY2QwODA3MzEwYzAxZGNjZmNlM2VlMjk1MTBkIl0sCiAgInByb2plY3RJZCI6ICJpY2NvLWNsb3VkIiwKICAibG9nc0J1Y2tldCI6ICJnczovLzk0MDM4MDE1NDYyMi5jbG91ZGJ1aWxkLWxvZ3MuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwKICAic291cmNlUHJvdmVuYW5jZSI6IHsKICB9LAogICJidWlsZFRyaWdnZXJJZCI6ICIzN2UyMWY0My0yNTMwLTQ1YzctYTU1Yy0xMDhiYTQwMzEwNGIiLAogICJvcHRpb25zIjogewogICAgInN1YnN0aXR1dGlvbk9wdGlvbiI6ICJBTExPV19MT09TRSIsCiAgICAibG9nZ2luZyI6ICJMRUdBQ1kiLAogICAgImR5bmFtaWNTdWJzdGl0dXRpb25zIjogdHJ1ZSwKICAgICJwb29sIjogewogICAgfQogIH0sCiAgImxvZ1VybCI6ICJodHRwczovL2NvbnNvbGUuY2xvdWQuZ29vZ2xlLmNvbS9jbG91ZC1idWlsZC9idWlsZHMvM2IzNGVhNDMtOTI4ZS00YjMyLTk1YzYtNDdmNDdjNjgzZDVlP3Byb2plY3RcdTAwM2Q5NDAzODAxNTQ2MjIiLAogICJzdWJzdGl0dXRpb25zIjogewogICAgIl9TRVJWSUNFX05BTUUiOiAibG9jYXRpdmUiLAogICAgIl9ERVBMT1lfUkVHSU9OIjogInVzLWNlbnRyYWwxIiwKICAgICJSRVZJU0lPTl9JRCI6ICI3ZWY3ZWU2NzUyMTVjY2QwODA3MzEwYzAxZGNjZmNlM2VlMjk1MTBkIiwKICAgICJUUklHR0VSX05BTUUiOiAibG9jYXRpdmUtZGVwbG95IiwKICAgICJfUExBVEZPUk0iOiAibWFuYWdlZCIsCiAgICAiX0lNQUdFX05BTUUiOiAiZ2NyLmlvL2ljY28tY2xvdWQvbG9jYXRpdmUuZ2FyZGVuIiwKICAgICJTSE9SVF9TSEEiOiAiN2VmN2VlNiIsCiAgICAiUkVQT19OQU1FIjogImxvY2F0aXZlLmdhcmRlbiIsCiAgICAiVFJJR0dFUl9CVUlMRF9DT05GSUdfUEFUSCI6ICIiLAogICAgIkNPTU1JVF9TSEEiOiAiN2VmN2VlNjc1MjE1Y2NkMDgwNzMxMGMwMWRjY2ZjZTNlZTI5NTEwZCIsCiAgICAiUkVGX05BTUUiOiAibWFpbiIsCiAgICAiX1RSSUdHRVJfSUQiOiAiMzdlMjFmNDMtMjUzMC00NWM3LWE1NWMtMTA4YmE0MDMxMDRiIiwKICAgICJCUkFOQ0hfTkFNRSI6ICJtYWluIgogIH0sCiAgInRhZ3MiOiBbImxvY2F0aXZlIiwgImRlcGxveSIsICJ0cmlnZ2VyLTM3ZTIxZjQzLTI1MzAtNDVjNy1hNTVjLTEwOGJhNDAzMTA0YiJdLAogICJ0aW1pbmciOiB7CiAgICAiRkVUQ0hTT1VSQ0UiOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0MDoyMS4xNjEwNDQ4NTBaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0MDoyNC4yNzg1MzY3NjBaIgogICAgfSwKICAgICJCVUlMRCI6IHsKICAgICAgInN0YXJ0VGltZSI6ICIyMDIyLTA0LTE2VDE1OjQwOjI0LjI3ODYzNDk4MFoiLAogICAgICAiZW5kVGltZSI6ICIyMDIyLTA0LTE2VDE1OjQ2OjIzLjE5NTg1Mjc3M1oiCiAgICB9LAogICAgIlBVU0giOiB7CiAgICAgICJzdGFydFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyMy4xOTU5NjQzMzBaIiwKICAgICAgImVuZFRpbWUiOiAiMjAyMi0wNC0xNlQxNTo0NjoyNi4wODg1Mjk0MjlaIgogICAgfQogIH0sCiAgImFydGlmYWN0cyI6IHsKICAgICJpbWFnZXMiOiBbImdjci5pby9pY2NvLWNsb3VkL2xvY2F0aXZlLmdhcmRlbjpsYXRlc3QiLCAiZ2NyLmlvL2ljY28tY2xvdWQvbG9jYXRpdmUuZ2FyZGVuOjdlZjdlZTY3NTIxNWNjZDA4MDczMTBjMDFkY2NmY2UzZWUyOTUxMGQiXQogIH0sCiAgInF1ZXVlVHRsIjogIjM2MDBzIiwKICAibmFtZSI6ICJwcm9qZWN0cy85NDAzODAxNTQ2MjIvbG9jYXRpb25zL2dsb2JhbC9idWlsZHMvM2IzNGVhNDMtOTI4ZS00YjMyLTk1YzYtNDdmNDdjNjgzZDVlIgp9","messageId":"4406496319062821","message_id":"4406496319062821","publishTime":"2022-04-16T15:46:27.233Z","publish_time":"2022-04-16T15:46:27.233Z"},"subscription":"projects/icco-cloud/subscriptions/builds"}`,
			Want: "GCB: SUCCESS [gcr.io/icco-cloud/locative.garden:latest gcr.io/icco-cloud/locative.garden:7ef7ee675215ccd0807310c01dccfce3ee29510d] @ <https://console.cloud.google.com/cloud-build/builds/3b34ea43-928e-4b32-95c6-47f47c683d5e?project=940380154622>",
		},
		"sonarr": {
			Have: `{ "episodes": [ { "id": 123, "episodeNumber": 1, "seasonNumber": 1, "title": "Test title", "qualityVersion": 0 } ], "eventType": "Test", "series": { "id": 1, "title": "Test Title", "path": "C:\\testpath", "tvdbId": 1234 } }`,
			Want: "Sonarr: Test Title 1x01 - \"Test\"\n",
		},
		"simple": {
			Have: `{"test":"bar","hi":"xyz"}`,
			Want: "hi: xyz\ntest: bar\n",
		},
		"lidarr": {
			Have: `{ "albums": [ { "id": 123, "title": "Test title", "qualityVersion": 0 } ], "eventType": "Test", "artist": { "id": 1, "name": "Test Name", "path": "C:\\testpath", "mbId": "aaaaa-aaa-aaaa-aaaaaa" } }`,
			Want: "Lidarr: Test Name - \"Test title\" - Test\n",
		},
		"plex TV Episode": {
			Have: `{"event":"media.resume","user":true,"owner":true,"Account":{"id":1,"thumb":"https://plex.tv/users/b921c1a7580bf543/avatar?c=1575240953","title":"icco"},"Server":{"title":"storm","uuid":"544b62f0b4f85d5d8f2c91696763d13578f5264a"},"Player":{"local":true,"publicAddress":"68.194.92.253","title":"BRAVIA 4K GB","uuid":"fb39bc1decfcc7a4-com-plexapp-android"},"Metadata":{"librarySectionType":"show","ratingKey":"56914","key":"/library/metadata/56914","parentRatingKey":"56851","grandparentRatingKey":"56850","guid":"com.plexapp.agents.thetvdb://301376/1/14?lang=en","parentGuid":"com.plexapp.agents.thetvdb://301376/1?lang=en","grandparentGuid":"com.plexapp.agents.thetvdb://301376?lang=en","type":"episode","title":"If You Can't Scooby-Doo the Time, Don't Scooby-Doo the Crime","grandparentTitle":"Be Cool, Scooby-Doo!","parentTitle":"Season 1","contentRating":"TV-G","summary":"Fred visits a high security prison and ends up investigating the ghost of an escaped inmate who's wreaking havoc on it.","index":14,"parentIndex":1,"rating":10.0,"lastViewedAt":1585962956,"year":2015,"thumb":"/library/metadata/56914/thumb/1585369990","art":"/library/metadata/56850/art/1585369991","parentThumb":"/library/metadata/56851/thumb/1585369991","grandparentThumb":"/library/metadata/56850/thumb/1585369991","grandparentArt":"/library/metadata/56850/art/1585369991","grandparentTheme":"/library/metadata/56850/theme/1585369991","originallyAvailableAt":"2015-10-31","addedAt":1585287583,"updatedAt":1585369990}}`,
		},
		"plex movie": {
			Have: `{ "event": "media.play", "user": true, "owner": true, "Account": { "id": 1, "thumb": "https://plex.tv/users/b921c1a7580bf543/avatar?c=1575240953", "title": "icco" }, "Server": { "title": "storm", "uuid": "544b62f0b4f85d5d8f2c91696763d13578f5264a" }, "Player": { "local": true, "publicAddress": "68.194.92.253", "title": "BRAVIA 4K GB", "uuid": "fb39bc1decfcc7a4-com-plexapp-android" }, "Metadata": { "librarySectionType": "movie", "ratingKey": "42593", "key": "/library/metadata/42593", "guid": "com.plexapp.agents.imdb://tt4154664?lang=en", "studio": "Marvel Studios", "type": "movie", "title": "Captain Marvel", "contentRating": "PG-13", "summary": "The story follows Carol Danvers as she becomes one of the universe’s most powerful heroes when Earth is caught in the middle of a galactic war between two alien races. Set in the 1990s, Captain Marvel is an all-new adventure from a previously unseen period in the history of the Marvel Cinematic Universe.", "rating": 7.8, "audienceRating": 4.8, "viewCount": 3, "lastViewedAt": 1562622295, "year": 2019, "tagline": "Higher. Further. Faster.", "thumb": "/library/metadata/42593/thumb/1587426457", "art": "/library/metadata/42593/art/1587426457", "duration": 7380000, "originallyAvailableAt": "2019-03-06", "addedAt": 1561930774, "updatedAt": 1587426457, "audienceRatingImage": "rottentomatoes://image.rating.spilled", "chapterSource": "media", "primaryExtraKey": "/library/metadata/42616", "ratingImage": "rottentomatoes://image.rating.ripe", "Genre": [ { "id": 146, "tag": "Action", "count": 455 }, { "id": 147, "tag": "Adventure", "count": 417 }, { "id": 149, "tag": "Sci-Fi", "count": 284 } ], "Director": [ { "id": 17785, "tag": "Ryan Fleck", "count": 3 }, { "id": 20354, "tag": "Anna Boden", "count": 2 } ], "Writer": [ { "id": 17786, "tag": "Ryan Fleck", "count": 3 }, { "id": 17787, "tag": "Anna Boden", "count": 3 }, { "id": 55136, "tag": "Geneva Robertson-Dworet", "count": 2 } ], "Producer": [ { "id": 4260, "tag": "Kevin Feige", "count": 23 } ], "Country": [ { "id": 52, "tag": "USA", "count": 968 } ], "Collection": [ { "id": 106698, "tag": "Captain Marvel" } ], "Role": [ { "id": 29541, "tag": "Brie Larson", "count": 8, "role": "Carol Danvers / Vers / Captain Marvel", "thumb": "http://image.tmdb.org/t/p/original/nTABNG8Sw4EVN3xodn2JBdEmBqv.jpg" }, { "id": 5006, "tag": "Samuel L. Jackson", "count": 38, "role": "Nick Fury", "thumb": "http://image.tmdb.org/t/p/original/mXN4Gw9tZJVKrLJHde2IcUHmV3P.jpg" }, { "id": 11107, "tag": "Ben Mendelsohn", "count": 10, "role": "Talos / Keller", "thumb": "http://image.tmdb.org/t/p/original/pA9mu9D2IolVA0v2Yo0tJm6uUyI.jpg" }, { "id": 2481, "tag": "Jude Law", "count": 15, "role": "Yon-Rogg", "thumb": "http://image.tmdb.org/t/p/original/xYXlyUh02Ue2PxYtkRbYAOKubb7.jpg" }, { "id": 3374, "tag": "Annette Bening", "count": 4, "role": "Supreme Intelligence / Dr. Wendy Lawson", "thumb": "http://image.tmdb.org/t/p/original/vVAvoiE6FQ4couqaB0ogaHR6Ef7.jpg" }, { "id": 10234, "tag": "Djimon Hounsou", "count": 11, "role": "Korath", "thumb": "http://image.tmdb.org/t/p/original/kC2AoZV3Wgtm854rEmaMt7YN2i.jpg" }, { "id": 17560, "tag": "Lee Pace", "count": 8, "role": "Ronan the Accuser", "thumb": "http://image.tmdb.org/t/p/original/8DVo5jbEmYpKPrhIFHkA7gGs1X8.jpg" }, { "id": 84822, "tag": "Lashana Lynch", "role": "Maria Rambeau", "thumb": "http://image.tmdb.org/t/p/original/eB4su7bV2ELijlDS9ZZyHSFlWkP.jpg" }, { "id": 14092, "tag": "Gemma Chan", "count": 6, "role": "Minn-Erva", "thumb": "http://image.tmdb.org/t/p/original/j8J5kZ4b9r0ByOYEfIqyjjgmu5s.jpg" }, { "id": 1514, "tag": "Clark Gregg", "count": 10, "role": "Agent Phil Coulson", "thumb": "http://image.tmdb.org/t/p/original/mq686D91XoZpqkzELn0888NOiZW.jpg" }, { "id": 84832, "tag": "Rune Temte", "role": "Bron-Charr" }, { "id": 84813, "tag": "Algenis Perez Soto", "role": "Att-Lass", "thumb": "http://image.tmdb.org/t/p/original/pUCNFbXpXuEelbrbiOo52vG3daa.jpg" }, { "id": 19352, "tag": "Mckenna Grace", "count": 3, "role": "Young Carol (13 Years Old)", "thumb": "http://image.tmdb.org/t/p/original/jQLBM6ErQnvU8QqNvW8KKF9y8N0.jpg" }, { "id": 84812, "tag": "Akira Akbar", "role": "Monica Rambeau (11 Years Old)", "thumb": "http://image.tmdb.org/t/p/original/zJ6IndgsiMHaM36jZfIBJZwzT5u.jpg" }, { "id": 16778, "tag": "Matthew Maher", "count": 5, "role": "Norex", "thumb": "http://image.tmdb.org/t/p/original/nO9Yhz7ERzATcHploljXgR4cAq.jpg" }, { "id": 84818, "tag": "Chuku Modu", "role": "Soh-Larr" }, { "id": 46143, "tag": "Vik Sahay", "count": 2, "role": "Hero Torfan", "thumb": "https://artworks.thetvdb.com/banners/actors/80673.jpg" }, { "id": 12451, "tag": "Colin Ford", "count": 3, "role": "Steve Danvers", "thumb": "http://image.tmdb.org/t/p/original/AquXBhH1jXdObhU1Fw3ecITEBVW.jpg" }, { "id": 39639, "tag": "Kenneth Mitchell", "count": 2, "role": "Joseph Danvers" }, { "id": 84834, "tag": "Stephen A. Chang", "role": "Cadet Officer", "thumb": "http://image.tmdb.org/t/p/original/gZxyFXDPDrZBAHqBXyujoYaW6Hn.jpg" }, { "id": 84829, "tag": "Pete Ploszek", "role": "Bret Johnson", "thumb": "http://image.tmdb.org/t/p/original/g0b5aePXb1IXgl7177BGDROe8IV.jpg" }, { "id": 84823, "tag": "London Fuller", "count": 2, "role": "Young Carol (6 Years Old)", "thumb": "http://image.tmdb.org/t/p/original/i8f3yFO82g2KXcl5CZySxcZcwlk.jpg" }, { "id": 84816, "tag": "Azari Akbar", "role": "Monica Rambeau (5 Years Old)", "thumb": "http://image.tmdb.org/t/p/original/2J424EHoWH1RQfAAxFN79FlbTi9.jpg" }, { "id": 84826, "tag": "Mark Daugherty", "role": "Skrull Main Tech" }, { "id": 84819, "tag": "Diana Toshiko", "role": "Skrull Tech #1", "thumb": "http://image.tmdb.org/t/p/original/pBQwfWbSUhtKH6gXQcHXJu7PS4e.jpg" }, { "id": 84817, "tag": "Barry Curtis", "role": "Mall Security Guard", "thumb": "http://image.tmdb.org/t/p/original/cvy6VCnSKxWO2K9NtjjNnTjQCFW.jpg" }, { "id": 84820, "tag": "Emily Ozrey", "role": "Surfer Girl Talos #1", "thumb": "http://image.tmdb.org/t/p/original/rr4pd5AHS17nB2s2lPQRBTKHupG.jpg" }, { "id": 84811, "tag": "Abigaille Ozrey", "role": "Surfer Girl Talos #2", "thumb": "http://image.tmdb.org/t/p/original/kTQtDO7lUL1ZqlWj3B57utALmLB.jpg" }, { "id": 84825, "tag": "Marilyn Brett", "role": "Older Lady on Train" }, { "id": 4257, "tag": "Stan Lee", "count": 29, "role": "Stan Lee", "thumb": "http://image.tmdb.org/t/p/original/nqMKapGZyCqpEqCbv3HTzxFeynY.jpg" }, { "id": 26292, "tag": "Robert Kazinsky", "count": 4, "role": "Biker (The Don)" }, { "id": 20611, "tag": "Nelson Franklin", "count": 4, "role": "Medical Examiner", "thumb": "http://image.tmdb.org/t/p/original/2YYnTxpFmS5i9qRkb6H4Nxky2px.jpg" }, { "id": 84828, "tag": "Patrick Brennan", "role": "Bartender" }, { "id": 31215, "tag": "Patrick Gallagher", "count": 3, "role": "Security Chief", "thumb": "http://image.tmdb.org/t/p/original/8ZWcp7kRwoh5cGZFByDQdhTODw9.jpg" }, { "id": 84814, "tag": "Ana Ayora", "role": "Agent Whitcher" }, { "id": 84824, "tag": "Lyonetta Flowers", "role": "Monica's Grandmother" }, { "id": 84831, "tag": "Rufus Flowers", "role": "Monica's Grandfather" }, { "id": 84833, "tag": "Sharon Blynn", "count": 2, "role": "Soren", "thumb": "http://image.tmdb.org/t/p/original/m3lzrdrNc0lleW9Tva64flZ8Fah.jpg" }, { "id": 84815, "tag": "Auden L. Ophuls", "role": "Talos' Daughter", "thumb": "http://image.tmdb.org/t/p/original/jJ6CqPbwg78k3NBlk9Z3QAwVjQs.jpg" }, { "id": 84821, "tag": "Harriet L. Ophuls", "role": "Talos' Daughter", "thumb": "http://image.tmdb.org/t/p/original/21hkUZE99Nff7G4PPr2jIPE2cnY.jpg" }, { "id": 84827, "tag": "Matthew Bellows", "role": "Accuser #1" }, { "id": 84830, "tag": "Richard Zeringue", "role": "Tom the Neighbor", "thumb": "http://image.tmdb.org/t/p/original/hWDoncFVG2PuE09b7rJ8fL5dvtH.jpg" }, { "id": 36069, "tag": "Duane Henry", "count": 2, "role": "Talos-Kree Soldier" }, { "id": 5003, "tag": "Chris Evans", "count": 16, "role": "Steve Rogers / Captain America (uncredited)", "thumb": "http://image.tmdb.org/t/p/original/7dUkkq1lK593XvOjunlUB11lKm1.jpg" }, { "id": 5004, "tag": "Scarlett Johansson", "count": 22, "role": "Natasha Romanoff / Black Widow (uncredited)", "thumb": "http://image.tmdb.org/t/p/original/6NsMbJXRlDZuDzatN2akFdGuTvx.jpg" }, { "id": 5007, "tag": "Don Cheadle", "count": 13, "role": "James 'Rhodey' Rhodes / War Machine (uncredited)", "thumb": "http://image.tmdb.org/t/p/original/b1EVJWdFn7a75qVYJgwO87W2TJU.jpg" }, { "id": 5002, "tag": "Mark Ruffalo", "count": 12, "role": "Bruce Banner / The Hulk (uncredited)", "thumb": "http://image.tmdb.org/t/p/original/z3dvKqMNDQWk3QLxzumloQVR0pv.jpg" }, { "id": 106695, "tag": "Matthew 'Spider' Kimmel", "role": "Spider" }, { "id": 106696, "tag": "Stephen 'Cajun' Del Bagno", "role": "Cajun" }, { "id": 106697, "tag": "Kelly Sue DeConnick", "role": "Lady in Subway Station (uncredited)" } ], "Similar": [ { "id": 84764, "tag": "Avengers: Endgame", "count": 13 }, { "id": 55754, "tag": "Ant-Man and the Wasp", "count": 26 }, { "id": 86837, "tag": "Spider-Man: Far from Home", "count": 21 }, { "id": 84766, "tag": "Shazam!", "count": 25 }, { "id": 76925, "tag": "Aquaman", "count": 14 }, { "id": 73692, "tag": "Spider-Man: Into the Spider-Verse", "count": 15 }, { "id": 6722, "tag": "Avengers: Infinity War", "count": 10 }, { "id": 6721, "tag": "Thor: Ragnarok", "count": 18 }, { "id": 62343, "tag": "Venom", "count": 17 }, { "id": 10107, "tag": "Black Panther", "count": 13 }, { "id": 85317, "tag": "Dark Phoenix", "count": 18 }, { "id": 76924, "tag": "Alita: Battle Angel", "count": 12 }, { "id": 6724, "tag": "Deadpool 2", "count": 15 }, { "id": 84835, "tag": "Pokémon Detective Pikachu", "count": 17 }, { "id": 73691, "tag": "Fantastic Beasts: The Crimes of Grindelwald", "count": 10 }, { "id": 5128, "tag": "Spider-Man: Homecoming", "count": 16 }, { "id": 33716, "tag": "Solo: A Star Wars Story", "count": 21 }, { "id": 76928, "tag": "Glass", "count": 9 }, { "id": 76930, "tag": "Ralph Breaks the Internet", "count": 16 }, { "id": 76926, "tag": "Bumblebee", "count": 10 } ] } }`,
		},
		"plex library add": {
			Have: `{"event":"library.new","user":true,"owner":true,"Account":{"id":1,"thumb":"https://plex.tv/users/b921c1a7580bf543/avatar?c=1575240953","title":"icco"},"Server":{"title":"storm","uuid":"544b62f0b4f85d5d8f2c91696763d13578f5264a"},"Metadata":{"librarySectionType":"show","ratingKey":"58657","key":"/library/metadata/58657/children","guid":"com.plexapp.agents.thetvdb://371040?lang=en","studio":"Tokyo MX","type":"show","title":"Wave, Listen to Me!","contentRating":"TV-14","summary":"The stage is Sapporo, Hokkaido. One night, our heroine, Minare Koda, spills her heartbroken woes to a radio station worker she meets while out drinking one night. The next day, she hears a recording of her pitiful grumbling being played live over the air. Minare storms into the station in a rage, only to then be duped by the station director into doing an impromptu talk show explaining her harsh dialogue. With just one recording, the many eccentric facets of Minare's life begin to pull every which direction as she falls ever deeper into the world of radio.","index":1,"year":2020,"thumb":"/library/metadata/58657/thumb/1587873096","art":"/library/metadata/58657/art/1587873096","duration":1500000,"originallyAvailableAt":"2020-04-04","leafCount":3,"viewedLeafCount":0,"childCount":1,"addedAt":1587872930,"updatedAt":1587873096,"Genre":[{"id":339,"tag":"Animation"},{"id":58093,"tag":"Anime"},{"id":48,"tag":"Comedy"},{"id":49,"tag":"Drama"},{"id":50,"tag":"Romance"}],"Role":[{"id":41911,"tag":"Kouki Uchiyama","role":"Shinji Oki","thumb":"https://artworks.thetvdb.com/banners/person/7884585/62109814.jpg"},{"id":112765,"tag":"Riho Sugiyama","role":"Minare Koda","thumb":"https://artworks.thetvdb.com/banners/person/8004244/62109759.jpg"},{"id":45876,"tag":"Sayaka Oohara","role":"Madoka Chishiro","thumb":"https://artworks.thetvdb.com/banners/person/7874988/62109785.jpg"},{"id":43344,"tag":"Kaito Ishikawa","role":"Ryuusuke Koumoto","thumb":"https://artworks.thetvdb.com/banners/person/303231/62109792.jpg"},{"id":6829,"tag":"Kazuhiro Yamaji","role":"Katsumi Kureko","thumb":"https://artworks.thetvdb.com/banners/person/307082/62109782.jpg"},{"id":112766,"tag":"Shinshuu Fuji","role":"Kanetsugu Matou","thumb":"https://artworks.thetvdb.com/banners/person/7878678/62109763.jpg"},{"id":112767,"tag":"Masaaki Yano","role":"Chuuya Nakahara","thumb":"https://artworks.thetvdb.com/banners/person/8160555/62109797.jpg"},{"id":47800,"tag":"Manaka Iwami","role":"Mizuho Nanba","thumb":"https://artworks.thetvdb.com/banners/person/465329/62109767.jpg"},{"id":46683,"tag":"Daisuke Namikawa","role":"Mitsuo Suga","thumb":"https://artworks.thetvdb.com/banners/person/292928/62109808.jpg"},{"id":41270,"tag":"Mamiko Noto","role":"Makie Tachibana","thumb":"https://artworks.thetvdb.com/banners/person/293451/62109800.jpg"},{"id":23621,"tag":"Bin Shimada","role":"Yoshiki Takarada","thumb":"https://artworks.thetvdb.com/banners/person/299647/62109805.jpg"}],"Location":[{"path":"/media/cask/TV/Wave, Listen to Me!"}]}}`,
			Want: "Plex: \"library.new\" - Wave, Listen to Me!\n",
		},
		"influx alert": {
			Have: `{"_check_id":"05f4817acbe02000","_check_name":"Load Test Check","_level":"ok","_measurement":"notifications","_message":"Check: Load Test Check is: ok","_notification_endpoint_id":"05f47c9c52d9b000","_notification_endpoint_name":"Relay","_notification_rule_id":"05f47cbfdb4b3000","_notification_rule_name":"Relay","_source_measurement":"system","_source_timestamp":1594004340000000000,"_start":"2020-07-06T02:49:00.289372603Z","_status_timestamp":1594004340165450419,"_stop":"2020-07-06T03:00:00.289372603Z","_time":"2020-07-06T03:00:00.289372603Z","_type":"threshold","_version":1,"host":"storm","load1":0.086}`,
			Want: "TICK Alert: \"Check: Load Test Check is: ok\"",
		},
		"gitlab push": {
			Have: `{"object_kind":"push","event_name":"push","before":"ef6835518ffe2a9b9e3aa9257b0140fb73672893","after":"ffa3543ef4fecf7b803595292c60be3f19f54b42","ref":"refs/heads/master","checkout_sha":"ffa3543ef4fecf7b803595292c60be3f19f54b42","message":null,"user_id":1514006,"user_name":"Nat Welch","user_username":"icco","user_email":"","user_avatar":"https://secure.gravatar.com/avatar/229e3746f6f5100c1d7d5d7a8a5b82d5?s=80\u0026d=identicon","project_id":10864002,"project":{"id":10864002,"name":"reliable-servers-with-go-book","description":"","web_url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book","avatar_url":null,"git_ssh_url":"git@gitlab.com:fullstackio/books/reliable-servers-with-go-book.git","git_http_url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book.git","namespace":"books","visibility_level":0,"path_with_namespace":"fullstackio/books/reliable-servers-with-go-book","default_branch":"master","ci_config_path":null,"homepage":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book","url":"git@gitlab.com:fullstackio/books/reliable-servers-with-go-book.git","ssh_url":"git@gitlab.com:fullstackio/books/reliable-servers-with-go-book.git","http_url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book.git"},"commits":[{"id":"ffa3543ef4fecf7b803595292c60be3f19f54b42","message":"Reword makeStorage and cleanupStorage sections, add benchmark thoughts\n","title":"Reword makeStorage and cleanupStorage sections, add benchmark thoughts","timestamp":"2020-09-14T10:21:49-04:00","url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/ffa3543ef4fecf7b803595292c60be3f19f54b42","author":{"name":"Steve McCarthy","email":"steve@redlua.com"},"added":[],"modified":["manuscript/chapters/testing-key-value-server.md","manuscript/resources/code/key-value-store/server_test.go"],"removed":[]},{"id":"cb1a5f9080de8a53d9d00bfa809f3621d98f4699","message":"writing\n","title":"writing","timestamp":"2020-09-14T13:55:10+00:00","url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/cb1a5f9080de8a53d9d00bfa809f3621d98f4699","author":{"name":"Nat Welch","email":"nat@natwelch.com"},"added":[],"modified":["manuscript/chapters/monitoring-key-value-server.md","manuscript/notes/chapter5-monitoring.md"],"removed":[]},{"id":"ef6835518ffe2a9b9e3aa9257b0140fb73672893","message":"Merge branch 'master' of gitlab.com:fullstackio/books/reliable-servers-with-go-book into master\n","title":"Merge branch 'master' of gitlab.com:fullstackio/books/reliable-servers-with-go-book into master","timestamp":"2020-09-13T14:26:51+00:00","url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/ef6835518ffe2a9b9e3aa9257b0140fb73672893","author":{"name":"Nat Welch","email":"nat@natwelch.com"},"added":[],"modified":["manuscript/chapters/monitoring-key-value-server.md","manuscript/chapters/testing-key-value-server.md","manuscript/resources/code/key-value-store/server_test.go"],"removed":[]}],"total_commits_count":3,"push_options":{},"repository":{"name":"reliable-servers-with-go-book","url":"git@gitlab.com:fullstackio/books/reliable-servers-with-go-book.git","description":"","homepage":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book","git_http_url":"https://gitlab.com/fullstackio/books/reliable-servers-with-go-book.git","git_ssh_url":"git@gitlab.com:fullstackio/books/reliable-servers-with-go-book.git","visibility_level":0}}`,
			Want: "GitLab push to reliable-servers-with-go-book: \n - \"Steve McCarthy\" @ 2020-09-14T10:21:49-04:00 https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/ffa3543ef4fecf7b803595292c60be3f19f54b42\n - \"Nat Welch\" @ 2020-09-14T13:55:10+00:00 https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/cb1a5f9080de8a53d9d00bfa809f3621d98f4699\n - \"Nat Welch\" @ 2020-09-13T14:26:51+00:00 https://gitlab.com/fullstackio/books/reliable-servers-with-go-book/-/commit/ef6835518ffe2a9b9e3aa9257b0140fb73672893\n",
		},
		"build working msg": {
			Have: `{"message":{"attributes":{"buildId":"04bf65f0-8152-4771-900f-17fd7757b607","status":"WORKING"},"data":"ewogICJpZCI6ICIwNGJmNjVmMC04MTUyLTQ3NzEtOTAwZi0xN2ZkNzc1N2I2MDciLAogICJzdGF0dXMiOiAiV09SS0lORyIsCiAgInNvdXJjZSI6IHsKICB9LAogICJjcmVhdGVUaW1lIjogIjIwMjEtMDQtMDNUMDM6MDU6NDMuOTY5NzMwOTI5WiIsCiAgInN0YXJ0VGltZSI6ICIyMDIxLTA0LTAzVDAzOjA1OjQ1LjM0MDk1OTk1MFoiLAogICJzdGVwcyI6IFt7CiAgICAibmFtZSI6ICJnY3IuaW8vY2xvdWQtYnVpbGRlcnMvZG9ja2VyIiwKICAgICJhcmdzIjogWyJidWlsZCIsICItdCIsICJnY3IuaW8vaWNjby1jbG91ZC9yZXBvcnRkOjQxNDk1OGEzMzYzYmNlMWJkZmJjNTUzYmUzYmJkN2VkZGRhZDFiODciLCAiLWYiLCAiRG9ja2VyZmlsZSIsICIuIl0KICB9XSwKICAidGltZW91dCI6ICI2MDBzIiwKICAiaW1hZ2VzIjogWyJnY3IuaW8vaWNjby1jbG91ZC9yZXBvcnRkOjQxNDk1OGEzMzYzYmNlMWJkZmJjNTUzYmUzYmJkN2VkZGRhZDFiODciXSwKICAicHJvamVjdElkIjogImljY28tY2xvdWQiLAogICJsb2dzQnVja2V0IjogImdzOi8vOTQwMzgwMTU0NjIyLmNsb3VkYnVpbGQtbG9ncy5nb29nbGV1c2VyY29udGVudC5jb20iLAogICJidWlsZFRyaWdnZXJJZCI6ICI1N2RkYzUzOC0wMmE3LTQ2ZjMtODYyMS02MjY4OTBjNDZlZGEiLAogICJvcHRpb25zIjogewogICAgInN1YnN0aXR1dGlvbk9wdGlvbiI6ICJBTExPV19MT09TRSIsCiAgICAibG9nZ2luZyI6ICJMRUdBQ1kiLAogICAgImR5bmFtaWNTdWJzdGl0dXRpb25zIjogdHJ1ZQogIH0sCiAgImxvZ1VybCI6ICJodHRwczovL2NvbnNvbGUuY2xvdWQuZ29vZ2xlLmNvbS9jbG91ZC1idWlsZC9idWlsZHMvMDRiZjY1ZjAtODE1Mi00NzcxLTkwMGYtMTdmZDc3NTdiNjA3P3Byb2plY3RcdTAwM2Q5NDAzODAxNTQ2MjIiLAogICJzdWJzdGl0dXRpb25zIjogewogICAgIlJFUE9fTkFNRSI6ICJyZXBvcnRkIiwKICAgICJSRVZJU0lPTl9JRCI6ICI0MTQ5NThhMzM2M2JjZTFiZGZiYzU1M2JlM2JiZDdlZGRkYWQxYjg3IiwKICAgICJDT01NSVRfU0hBIjogIjQxNDk1OGEzMzYzYmNlMWJkZmJjNTUzYmUzYmJkN2VkZGRhZDFiODciLAogICAgIlNIT1JUX1NIQSI6ICI0MTQ5NThhIiwKICAgICJCUkFOQ0hfTkFNRSI6ICJtYXN0ZXIiCiAgfSwKICAidGFncyI6IFsidHJpZ2dlci01N2RkYzUzOC0wMmE3LTQ2ZjMtODYyMS02MjY4OTBjNDZlZGEiXSwKICAiYXJ0aWZhY3RzIjogewogICAgImltYWdlcyI6IFsiZ2NyLmlvL2ljY28tY2xvdWQvcmVwb3J0ZDo0MTQ5NThhMzM2M2JjZTFiZGZiYzU1M2JlM2JiZDdlZGRkYWQxYjg3Il0KICB9LAogICJxdWV1ZVR0bCI6ICIzNjAwcyIsCiAgIm5hbWUiOiAicHJvamVjdHMvOTQwMzgwMTU0NjIyL2xvY2F0aW9ucy9nbG9iYWwvYnVpbGRzLzA0YmY2NWYwLTgxNTItNDc3MS05MDBmLTE3ZmQ3NzU3YjYwNyIKfQ==","messageId":"2229909757575397","message_id":"2229909757575397","publishTime":"2021-04-03T03:05:46.886Z","publish_time":"2021-04-03T03:05:46.886Z"},"subscription":"projects/icco-cloud/subscriptions/builds"}`,
			Want: "",
		},
		"plex album": {
			Have: `{"event":"library.new","user":true,"owner":true,"Account":{"id":639131,"thumb":"https://plex.tv/users/b921c1a7580bf543/avatar?c=1662304482","title":"icco"},"Server":{"title":"storm","uuid":"544b62f0b4f85d5d8f2c91696763d13578f5264a"},"Metadata":{"librarySectionType":"artist","ratingKey":"239918","key":"/library/metadata/239918/children","parentRatingKey":"220417","guid":"local://239918","parentGuid":"plex://artist/5d07bbfd403c6402904a6579","type":"album","title":"Grave Diggers: Tom Waits","parentKey":"/library/metadata/220417","librarySectionTitle":"Music","librarySectionID":3,"librarySectionKey":"/library/sections/3","parentTitle":"Tom Waits","summary":"","index":1,"year":2022,"thumb":"/library/metadata/239918/thumb/1684767305","art":"/library/metadata/220417/art/1684748132","parentThumb":"/library/metadata/220417/thumb/1684748132","originallyAvailableAt":"2022-01-01","leafCount":6,"viewedLeafCount":0,"addedAt":1684767304,"updatedAt":1684767305,"Genre":[{"id":493305,"filter":"genre=493305","tag":"Rock"}]}}`,
			Want: "Plex: \"library.new\" - Grave Diggers: Tom Waits by Tom Waits\n",
		},
		"radarr": {
			Have: `{ "movie": { "id": 2403, "title": "Blockers", "year": 2018, "releaseDate": "2018-07-03", "folderPath": "/data/Movies/Blockers (2018)", "tmdbId": 437557, "imdbId": "tt2531344", "overview": "When three parents discover that each of their daughters have a pact to lose their virginity at prom, they launch a covert one-night operation to stop the teens from sealing the deal." }, "remoteMovie": { "tmdbId": 437557, "imdbId": "tt2531344", "title": "Blockers", "year": 2018 }, "movieFile": { "id": 2850, "relativePath": "Blockers (2018) Bluray-1080p.mkv", "path": "/data/Downloads/complete/Blockers.2018.BluRay.1080p.DTS-HD.MA.5.1.x264-LEGi0N-AsRequested/Blockers.2018.BluRay.1080p.DTS-HD.MA.5.1.x264-LEGi0N-AsRequested.mkv", "quality": "Bluray-1080p", "qualityVersion": 1, "releaseGroup": "LEGi0N", "sceneName": "Blockers.2018.BluRay.1080p.DTS-HD.MA.5.1.x264-LEGi0N-AsRequested", "indexerFlags": "0", "size": 12568531076, "dateAdded": "2023-07-03T19:28:32.741575Z", "mediaInfo": { "audioChannels": 5.1, "audioCodec": "DTS-HD MA", "audioLanguages": [ "eng" ], "height": 800, "width": 1920, "subtitles": [ "eng" ], "videoCodec": "x264", "videoDynamicRange": "", "videoDynamicRangeType": "" } }, "isUpgrade": false, "downloadClient": "sabnzbd", "downloadClientType": "SABnzbd", "downloadId": "SABnzbd_nzo_spb4z0dx", "customFormatInfo": { "customFormats": [], "customFormatScore": 0 }, "release": { "releaseTitle": "Blockers.2018.BluRay.1080p.DTS-HD.MA.5.1.x264-LEGi0N-AsRequested", "indexer": "nzbgeek", "size": 13749100000 }, "eventType": "Download", "instanceName": "Radarr", "applicationUrl": "" }`,
			Want: "",
		},
		"readarr": {
			Have: `{ "author": { "id": 157, "name": "Steven Brust", "path": "/data/Books/Steven Brust", "goodreadsId": "27704" }, "books": [ { "id": 23128, "goodreadsId": "95912760", "title": "Tsalmoth", "edition": { "title": "Tsalmoth", "goodreadsId": "60820532", "asin": "B09XL5FXZ3" }, "releaseDate": "2023-04-25T04:00:00Z" } ], "release": { "quality": "EPUB", "qualityVersion": 1, "releaseTitle": "Steven Brust - Tsalmoth (epub)", "indexer": "nzbgeek", "size": 592000, "customFormatScore": 0, "customFormats": [] }, "downloadClient": "sabnzbd", "downloadClientType": "SABnzbd", "downloadId": "SABnzbd_nzo_fjf_tdta", "eventType": "Grab", "instanceName": "Readarr" }`,
			Want: "Readarr: Steven Brust - \"Tsalmoth\" - Grab\n",
		},
		"update": {
			Have: `{"deployed": "relay", "image": "gcr.io/icco-cloud/relay:0fd7bb0ecd170b145417563156e8ab77eb265f9e"}`,
			Want: `Deployed: "relay" -> "gcr.io/icco-cloud/relay:0fd7bb0ecd170b145417563156e8ab77eb265f9e"`,
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
