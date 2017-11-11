package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// App represents a Veracode Application Profile
type App struct {
	AppID   string `xml:"app_id,attr"`
	AppName string `xml:"app_name,attr"`
}

// ParseAppList calls the Veracode getapplist.do API and returns an array of Apps
func ParseAppList(credsFile string) ([]App, error) {
	var apps []App

	appListAPI, err := appList(credsFile)
	if err != nil {
		return nil, err
	}
	decoder := xml.NewDecoder(bytes.NewReader(appListAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for flaw
			if se.Name.Local == "app" {
				var app App
				decoder.DecodeElement(&app, &se)
				apps = append(apps, app)
			}
			if se.Name.Local == "error" {
				return nil, errors.New("api for GetAppList returned with an error element")
			}
		}
	}
	return apps, nil
}
