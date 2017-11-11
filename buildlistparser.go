package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// A Build represents a Veracode Build within an application.
type Build struct {
	BuildID           string `xml:"build_id,attr"`
	Version           string `xml:"version,attr"`
	PolicyUpdatedDate string `xml:"policy_updated_date,attr"`
}

// ParseBuildList calls the Veracode getbuildlist.do API and returns an array of Builds
func ParseBuildList(credsFile, appID string) ([]Build, error) {
	var builds []Build

	buildListAPI, err := buildList(credsFile, appID)
	if err != nil {
		return nil, err
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
				return nil, errors.New("api for GetBuildList returned with an error element")
			}
		}
	}
	return builds, nil
}
