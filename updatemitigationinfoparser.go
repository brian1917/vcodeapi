package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// ParseUpdateMitigation process an update mitigation request and returns an error if applicable
func ParseUpdateMitigation(credsFile, buildID, action, comment, flawList string) error {

	updateMitigationAPI, err := updateMitigationInfo(credsFile, buildID, action, comment, flawList)
	if err != nil {
		return err
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
				return fmt.Errorf("updatemitigationinfo.do error element returned when updating mitigation info for flaw IDs %v in build ID %v",
					flawList, buildID)
			}
		}
	}
	return nil
}
