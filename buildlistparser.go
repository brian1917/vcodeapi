package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

type Build struct {
	BuildID           string `xml:"build_id,attr"`
	Version           string `xml:"version,attr"`
	PolicyUpdatedDate string `xml:"policy_updated_date,attr"`
}

func ParseBuildList(username, password, app_id string) ([]Build, error) {
	var builds []Build
	var errMsg error = nil

	buildListAPI, err := buildList(username, password, app_id)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(buildListAPI))
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
			if se.Name.Local == "build" {
				var build Build
				decoder.DecodeElement(&build, &se)
				builds = append(builds, build)
			}
			if se.Name.Local == "error" {
				errMsg = errors.New("api for GetBuildList returned with an error element")
			}
		}
	}
	return builds, errMsg
}
