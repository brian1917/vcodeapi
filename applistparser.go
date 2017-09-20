package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

type App struct {
	AppID   string `xml:"app_id,attr"`
	AppName string `xml:"app_name,attr"`
}

func ParseAppList(username, password string) ([]App, error) {
	var apps []App
	var errMsg error = nil

	appListAPI, err := appList(username, password)
	if err!= nil{
		log.Fatal(err)
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
				errMsg = errors.New("api for GetAppList returned with an error element")
			}
		}
	}
	return apps, errMsg
}
