package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"
)

// ParseUpdateMitigation process an update mitigation request and returns an error if applicable
func ParseUpdateMitigation(credsFile, buildID, action, comment, flawList string) error {
	var errMsg error

	updateMitigationAPI, err := updateMitigationInfo(credsFile, buildID, action, comment, flawList)
	if err != nil {
		log.Fatal(err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(updateMitigationAPI))
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
			if se.Name.Local == "error" {
				errMsg = errors.New("api for UpdateMitigationInfo returned with an error element")
			}
		}
	}
	return errMsg
}
