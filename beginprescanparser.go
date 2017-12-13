package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// ParseBeginPreScan process an begin prescan request and returns an error if applicable
func ParseBeginPreScan(credsFile, appID, sandboxID string, autoScan bool) error {

	beginPreScanAPI, err := beginPreScan(credsFile, appID, sandboxID, autoScan)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(bytes.NewReader(beginPreScanAPI))
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
				return fmt.Errorf("beginprescan.do error element returned")
			}
		}
	}
	return nil
}
