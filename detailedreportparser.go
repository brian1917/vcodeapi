package vcodeapi

import (
	"bytes"
	"encoding/xml"
	"errors"
)

type detReport struct {
	AppName                string `xml:"app_name,attr"`
	AppID                  string `xml:"app_id,attr"`
	PolicyName             string `xml:"policy_name,attr"`
	PolicyComplianceStatus string `xml:"policy_compliance_status,attr"`
	PolicyRulesStatus      string `xml:"policy_rules_status,attr"`
	GracePeriodExpired     string `xml:"grace_period_expired,attr"`
	BusinessUnit           string `xml:"business_unit,attr"`
}

// Flaw represents a finding from a Veracode test (static, dynamic, or MPT)
type Flaw struct {
	Issueid                 string `xml:"issueid,attr"`
	CweName                 string `xml:"categoryname,attr"`
	CategoryID              string `xml:"categoryid,attr"`
	CategoryName            string
	Cweid                   string `xml:"cweid,attr"`
	RemediationStatus       string `xml:"remediation_status,attr"`
	MitigationStatus        string `xml:"mitigation_status,attr"`
	AffectsPolicyCompliance string `xml:"affects_policy_compliance,attr"`
	PolicyName              string
	DateFirstOccurrence     string      `xml:"date_first_occurrence,attr"`
	Severity                string      `xml:"severity,attr"`
	ExploitLevel            string      `xml:"exploitLevel,attr"`
	Module                  string      `xml:"module,attr"`
	Sourcefile              string      `xml:"sourcefile,attr"`
	Line                    string      `xml:"line,attr"`
	Description             string      `xml:"description,attr"`
	Mitigations             Mitigations `xml:"mitigations"`
	Annotations             Annotations `xml:"annotations"`
}

// Mitigations are an array individual mitigations
type Mitigations struct {
	Mitigation []Mitigation `xml:"mitigation"`
}

// Mitigation is an individual documentation of a compensating control or reason a policy-violating flaw will not be addressed.
type Mitigation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

// Annotations are an array of individual annotations (comments)
type Annotations struct {
	Annotation []Annotation `xml:"annotation"`
}

// Annotation is a comment on a flaw (separate from comments attached to mitigation actions)
type Annotation struct {
	Action      string `xml:"action,attr"`
	Description string `xml:"description,attr"`
	User        string `xml:"user,attr"`
	Date        string `xml:"date,attr"`
}

// CustomField is metadata for an application profile (extracted from detailed report API)
type CustomField struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// ParseDetailedReport parses the detailedreport.do API and returns an array of Flaws and Custom Fields.
func ParseDetailedReport(credsFile, buildID string) ([]Flaw, []CustomField, error) {
	var flaws []Flaw
	var customFields []CustomField
	var detRep detReport

	detailedReportAPI, err := detailedReport(credsFile, buildID)
	if err != nil {
		return nil, nil, err
	}

	/**
	Two decoders isn't ideal, but we are avoiding putting entire document into structs.
	We are decoding the entire document in decoder1 to get metadeta around the build.
	We decode flaws and custom fields in decoder 2.
	**/

	//decoder1 gets information on the app
	decoder1 := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder1.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "detailedreport" {
				decoder1.DecodeElement(&detRep, &se)
			}
		}
	}

	//decoder 2 gets information on flaws and custom fields.
	decoder2 := xml.NewDecoder(bytes.NewReader(detailedReportAPI))
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder2.Token()

		if t == nil {
			break
		}
		// Inspect the type of the token just read
		switch se := t.(type) {
		case xml.StartElement:
			// Read StartElement and check for errors, flaws, and custom field
			if se.Name.Local == "error" {
				return nil, nil, errors.New("api for GetDetailedReport returned with an error element")
			}
			if se.Name.Local == "flaw" {
				var flaw Flaw
				decoder2.DecodeElement(&flaw, &se)
				flaw.CategoryName = categoryMap[flaw.CategoryID]
				flaw.PolicyName = detRep.PolicyName
				flaws = append(flaws, flaw)
			}
			if se.Name.Local == "customfield" {
				var cField CustomField
				decoder2.DecodeElement(&cField, &se)
				customFields = append(customFields, cField)
			}
		}
	}
	return flaws, customFields, nil

}
