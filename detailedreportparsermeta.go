package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"log"
)

// DetReportMeta represents metadata in the detailed report XML
type DetReportMeta struct {
	AppName                string `xml:"app_name,attr"`
	AppID                  string `xml:"app_id,attr"`
	PolicyName             string `xml:"policy_name,attr"`
	PolicyComplianceStatus string `xml:"policy_compliance_status,attr"`
	PolicyRulesStatus      string `xml:"policy_rules_status,attr"`
	GracePeriodExpired     string `xml:"grace_period_expired,attr"`
	BusinessUnit           string `xml:"business_unit,attr"`
}

// ParseBuildMetaData parses the detailedreport.do API and returns a detailed report meta data struct.
func ParseBuildMetaData(credsFile, buildID string) (DetReportMeta, error) {
	var errMsg error
	var detReportMeta DetReportMeta

	detailedReportAPI, err := detailedReport(credsFile, buildID)
	if err != nil {
		log.Fatal(err)
	}

	//decoder1 gets information on the app
	decoder := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "detailedreport" {
				decoder.DecodeElement(&detReportMeta, &se)
			}
		}
	}

	return detReportMeta, errMsg

}
