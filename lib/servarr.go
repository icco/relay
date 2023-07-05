package lib

type Servarr struct {
	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
	EventType       string `json:"eventType"`
	InstanceName    string `json:"instanceName"`
}

//  "eventType": "ApplicationUpdate",
