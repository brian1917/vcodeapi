package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// ParseUploadFile processes a file upload request and returns an error if applicable
func ParseUploadFile(credsFile, appID, sandboxID, path string) error {

	uploadFileAPI, err := uploadFile(credsFile, appID, sandboxID, path)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(bytes.NewReader(uploadFileAPI))
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
				return fmt.Errorf("uploadfile.do error element returned")
			}
		}
	}
	return nil
}
